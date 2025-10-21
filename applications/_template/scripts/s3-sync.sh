#!/bin/bash
# S3 sync helper script
# Utilities for syncing data to/from S3 with progress tracking

set -euo pipefail

# Function: Show usage
show_usage() {
    cat <<EOF
Usage: s3-sync.sh [COMMAND] [OPTIONS]

Commands:
    download    Download data from S3
    upload      Upload data to S3
    list        List S3 bucket contents

Download Options:
    --bucket BUCKET     S3 bucket name
    --prefix PREFIX     S3 prefix/path
    --dest DIR          Local destination directory
    --parallel N        Number of parallel transfers (default: 10)

Upload Options:
    --source DIR        Local source directory
    --bucket BUCKET     S3 bucket name
    --prefix PREFIX     S3 prefix/path
    --parallel N        Number of parallel transfers (default: 10)

List Options:
    --bucket BUCKET     S3 bucket name
    --prefix PREFIX     S3 prefix/path (optional)

Examples:
    # Download input data
    s3-sync.sh download \\
        --bucket my-data \\
        --prefix input/2024/ \\
        --dest /data/input

    # Upload results
    s3-sync.sh upload \\
        --source /data/output \\
        --bucket my-results \\
        --prefix run-001/

    # List bucket contents
    s3-sync.sh list --bucket my-data --prefix input/

EOF
}

# Function: Download from S3
s3_download() {
    local bucket=""
    local prefix=""
    local dest=""
    local parallel=10

    while [[ $# -gt 0 ]]; do
        case $1 in
            --bucket) bucket="$2"; shift 2 ;;
            --prefix) prefix="$2"; shift 2 ;;
            --dest) dest="$2"; shift 2 ;;
            --parallel) parallel="$2"; shift 2 ;;
            *) echo "Unknown option: $1"; return 1 ;;
        esac
    done

    if [[ -z "$bucket" ]] || [[ -z "$dest" ]]; then
        echo "Error: --bucket and --dest are required"
        return 1
    fi

    local s3_path="s3://${bucket}"
    if [[ -n "$prefix" ]]; then
        s3_path="${s3_path}/${prefix}"
    fi

    echo "Downloading from ${s3_path} to ${dest}..."
    mkdir -p "$dest"

    local start_time=$(date +%s)

    # Use aws s3 sync with progress
    aws s3 sync "$s3_path" "$dest" \
        --no-progress \
        --only-show-errors

    local end_time=$(date +%s)
    local duration=$((end_time - start_time))
    local size=$(du -sh "$dest" | cut -f1)

    echo "Downloaded ${size} in ${duration} seconds"
}

# Function: Upload to S3
s3_upload() {
    local source=""
    local bucket=""
    local prefix=""
    local parallel=10

    while [[ $# -gt 0 ]]; do
        case $1 in
            --source) source="$2"; shift 2 ;;
            --bucket) bucket="$2"; shift 2 ;;
            --prefix) prefix="$2"; shift 2 ;;
            --parallel) parallel="$2"; shift 2 ;;
            *) echo "Unknown option: $1"; return 1 ;;
        esac
    done

    if [[ -z "$source" ]] || [[ -z "$bucket" ]]; then
        echo "Error: --source and --bucket are required"
        return 1
    fi

    if [[ ! -d "$source" ]]; then
        echo "Error: Source directory does not exist: ${source}"
        return 1
    fi

    local s3_path="s3://${bucket}"
    if [[ -n "$prefix" ]]; then
        s3_path="${s3_path}/${prefix}"
    fi

    echo "Uploading from ${source} to ${s3_path}..."

    local start_time=$(date +%s)
    local size=$(du -sh "$source" | cut -f1)

    # Use aws s3 sync with progress
    aws s3 sync "$source" "$s3_path" \
        --no-progress \
        --only-show-errors

    local end_time=$(date +%s)
    local duration=$((end_time - start_time))

    echo "Uploaded ${size} in ${duration} seconds"
}

# Function: List S3 bucket contents
s3_list() {
    local bucket=""
    local prefix=""

    while [[ $# -gt 0 ]]; do
        case $1 in
            --bucket) bucket="$2"; shift 2 ;;
            --prefix) prefix="$2"; shift 2 ;;
            *) echo "Unknown option: $1"; return 1 ;;
        esac
    done

    if [[ -z "$bucket" ]]; then
        echo "Error: --bucket is required"
        return 1
    fi

    local s3_path="s3://${bucket}"
    if [[ -n "$prefix" ]]; then
        s3_path="${s3_path}/${prefix}"
    fi

    echo "Listing contents of ${s3_path}..."
    aws s3 ls "$s3_path" --recursive --human-readable --summarize
}

# Main script logic
if [[ $# -eq 0 ]]; then
    show_usage
    exit 0
fi

COMMAND="$1"
shift

case $COMMAND in
    download)
        s3_download "$@"
        ;;
    upload)
        s3_upload "$@"
        ;;
    list)
        s3_list "$@"
        ;;
    --help|-h)
        show_usage
        exit 0
        ;;
    *)
        echo "Error: Unknown command: $COMMAND"
        show_usage
        exit 1
        ;;
esac
