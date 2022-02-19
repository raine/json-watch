assert_file_equals() {
  local -r file="$1"
  local -r contents="$2"
  if [ "$(cat "$file")" != "$contents" ]; then
    local -r rem="$BATSLIB_FILE_PATH_REM"
    local -r add="$BATSLIB_FILE_PATH_ADD"
    batslib_print_kv_single 4 'path' "${file/$rem/$add}" \
      | batslib_decorate 'file does not equal to expected' \
      | fail
  fi
}
