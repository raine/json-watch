setup() {
  load 'test_helper/bats-support/load'
  load 'test_helper/bats-assert/load'
  load 'test_helper/bats-file/load'
  load 'test_helper/custom'
  DIR="$( cd "$( dirname "$BATS_TEST_FILENAME" )" >/dev/null 2>&1 && pwd )"
  PATH="$DIR/../bin:$PATH"
  config="${BATS_FILE_TMPDIR}json-watch"
  JSON_WATCH_CONFIG_DIR="$config"
  export JSON_WATCH_CONFIG_DIR config
  rm -rf "$config"
}

diag() {
  echo "# DEBUG $@" >&3
}

@test "no args shows help" {
  run json-watch
  assert_output -p "Usage: cat data.json | json-watch <name>"
  assert_output -p "Error: name is required"
  assert_failure
}

@test "key is required" {
  run json-watch test
  assert_output -p "Error: key is required"
  assert_failure
}

@test "first run with empty watch file" {
  run bash -c 'echo '"'"'[{"id":1},{"id":2}]'"'"' | json-watch test --key id'
  assert_success
  assert_output ""
  assert_file_equals "$config/watches/test" $'1\n2'
}

@test "second run outputs unseen objects and writes id to watchfile" {
  run bash -c 'echo '"'"'[{"id":1},{"id":2}]'"'"' | json-watch test --key id'
  run bash -c 'echo '"'"'[{"id":1},{"id":2},{"id":3},{"id":4}]'"'"' | json-watch test --key id'
  assert_output $'{"id":3}\n{"id":4}'
  assert_file_equals "$config/watches/test" $'1\n2\n3\n4'
}

@test "string property as key works" {
  run bash -c 'echo '"'"'[{"id":"1"}]'"'"' | json-watch test --key id'
  assert_file_equals "$config/watches/test" '1'
}

@test "boolean property as key errors" {
  run bash -c 'echo '"'"'[{"id":true}]'"'"' | json-watch test --key id'
  assert_failure
  assert_output "Error: Got an object with the property 'id' with value that is not a string or number"
}

@test "object as root value in stdin errors" {
  run bash -c 'echo '"'"'{"id":1}'"'"' | json-watch test --key id'
  assert_failure
  assert_output "Error: Expecting array of objects in stdin"
}
