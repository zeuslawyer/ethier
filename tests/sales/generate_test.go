package sales_test

//go:generate sh -c "solc *.sol --base-path ../../ --include-path ../../node_modules --combined-json abi,bin | abigen --combined-json /dev/stdin --pkg sales_test --out generated_test.go"
