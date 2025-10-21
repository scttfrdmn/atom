#!/bin/bash
# Application entrypoint script
# Handles setup, execution, and cleanup for application runs

set -euo pipefail

# Source HPC environment
source /etc/profile.d/hpc-base.sh

# Print environment information
echo "=========================================="
echo "Application: your-app"
echo "Version: ${APP_VERSION}"
echo "Architecture: $(uname -m)"
echo "Date: $(date)"
echo "=========================================="
echo ""

# Function: Show usage
show_usage() {
    cat <<EOF
Usage: docker run your-app [OPTIONS]

Options:
    --help                  Show this help message
    --version               Show application version
    --input S3_PATH         S3 path to input data (required)
    --output S3_PATH        S3 path for output data (required)
    --config FILE           Configuration file (optional)
    --param KEY=VALUE       Override configuration parameter

Examples:
    # Run with S3 input/output
    docker run your-app \\
        --input s3://bucket/input/ \\
        --output s3://bucket/output/

    # Run with custom config
    docker run your-app \\
        --input s3://bucket/input/ \\
        --output s3://bucket/output/ \\
        --config /data/custom-config.yaml

Environment Variables:
    APP_DATA         Input data directory (default: /data/input)
    APP_OUTPUT       Output data directory (default: /data/output)
    OMP_NUM_THREADS  Number of OpenMP threads (default: auto-detect)

EOF
}

# Function: Show version
show_version() {
    echo "your-app version ${APP_VERSION}"
    echo "Platform: $(uname -s) $(uname -m)"
    echo "Compiler: $(gcc --version | head -n1)"
}

# Parse command-line arguments
INPUT_S3=""
OUTPUT_S3=""
CONFIG_FILE=""
declare -A PARAMS

while [[ $# -gt 0 ]]; do
    case $1 in
        --help)
            show_usage
            exit 0
            ;;
        --version)
            show_version
            exit 0
            ;;
        --input)
            INPUT_S3="$2"
            shift 2
            ;;
        --output)
            OUTPUT_S3="$2"
            shift 2
            ;;
        --config)
            CONFIG_FILE="$2"
            shift 2
            ;;
        --param)
            IFS='=' read -r key value <<< "$2"
            PARAMS[$key]="$value"
            shift 2
            ;;
        *)
            echo "Error: Unknown option $1"
            show_usage
            exit 1
            ;;
    esac
done

# Validate required arguments
if [[ -z "$INPUT_S3" ]]; then
    echo "Error: --input is required"
    show_usage
    exit 1
fi

if [[ -z "$OUTPUT_S3" ]]; then
    echo "Error: --output is required"
    show_usage
    exit 1
fi

# Set OpenMP threads if not already set
if [[ -z "${OMP_NUM_THREADS:-}" ]]; then
    export OMP_NUM_THREADS=$(nproc)
    echo "Setting OMP_NUM_THREADS=${OMP_NUM_THREADS}"
fi

# Download input data from S3
echo "Downloading input data from ${INPUT_S3}..."
START_TIME=$(date +%s)

if [[ "$INPUT_S3" == s3://* ]]; then
    aws s3 sync "$INPUT_S3" "$APP_DATA/" --quiet
    echo "Downloaded input data ($(du -sh $APP_DATA | cut -f1))"
else
    echo "Error: Input path must be S3 URI (s3://...)"
    exit 1
fi

DOWNLOAD_TIME=$(($(date +%s) - START_TIME))
echo "Download completed in ${DOWNLOAD_TIME} seconds"
echo ""

# Generate or copy configuration
echo "Setting up configuration..."

if [[ -n "$CONFIG_FILE" ]]; then
    # Use provided config
    cp "$CONFIG_FILE" /opt/run-dir/config.yaml
    echo "Using configuration from ${CONFIG_FILE}"
else
    # Generate default config from template
    /app/config-generator \
        --template /app/templates/config.yaml.template \
        --output /opt/run-dir/config.yaml
    echo "Generated default configuration"
fi

# Apply parameter overrides
for key in "${!PARAMS[@]}"; do
    value="${PARAMS[$key]}"
    echo "Overriding parameter: ${key}=${value}"
    # Use sed or yq to modify config.yaml
    # This is a placeholder - adjust based on your config format
done

echo ""

# Run the application
echo "Starting application..."
echo "Working directory: $(pwd)"
echo "Input data: ${APP_DATA}"
echo "Output directory: ${APP_OUTPUT}"
echo ""

RUN_START_TIME=$(date +%s)

# Execute your application
# Adjust this command for your specific application
${APP_ROOT}/bin/your-app \
    --input "${APP_DATA}" \
    --output "${APP_OUTPUT}" \
    --config /opt/run-dir/config.yaml \
    2>&1 | tee /opt/run-dir/application.log

EXIT_CODE=${PIPESTATUS[0]}

RUN_TIME=$(($(date +%s) - RUN_START_TIME))
echo ""
echo "Application completed in ${RUN_TIME} seconds"

if [[ $EXIT_CODE -ne 0 ]]; then
    echo "Error: Application exited with code ${EXIT_CODE}"
    echo "Check /opt/run-dir/application.log for details"
    exit $EXIT_CODE
fi

# Upload results to S3
echo ""
echo "Uploading results to ${OUTPUT_S3}..."
UPLOAD_START_TIME=$(date +%s)

if [[ "$OUTPUT_S3" == s3://* ]]; then
    # Also upload the log file
    cp /opt/run-dir/application.log "${APP_OUTPUT}/"

    aws s3 sync "${APP_OUTPUT}/" "$OUTPUT_S3" --quiet
    echo "Uploaded results ($(du -sh $APP_OUTPUT | cut -f1))"
else
    echo "Error: Output path must be S3 URI (s3://...)"
    exit 1
fi

UPLOAD_TIME=$(($(date +%s) - UPLOAD_START_TIME))
echo "Upload completed in ${UPLOAD_TIME} seconds"

# Print summary
TOTAL_TIME=$(($(date +%s) - START_TIME))
echo ""
echo "=========================================="
echo "Execution Summary"
echo "=========================================="
echo "Download time:    ${DOWNLOAD_TIME}s"
echo "Computation time: ${RUN_TIME}s"
echo "Upload time:      ${UPLOAD_TIME}s"
echo "Total time:       ${TOTAL_TIME}s"
echo "=========================================="
echo ""
echo "Results available at: ${OUTPUT_S3}"

exit 0
