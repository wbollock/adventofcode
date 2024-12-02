# justfile

set shell := ["bash", "-eu", "-o", "pipefail", "-c"]

# Command to create a new day of files
new-day:
    @last_day=$(ls 2024/input/day* | xargs -n1 basename | grep -oE '[0-9]+' | sort -n | tail -n1) && \
    new_day=$(printf "%02d" $((10#$last_day + 1))) && \
    touch 2024/input/day$new_day && \
    touch 2024/input/day$new_day"test" && \
    touch 2024/solution/day$new_day".go" && \
    touch 2024/solution/day$new_day"_test.go"