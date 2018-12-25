abigen:
	go build github.com/ethereum/go-ethereum/cmd/abigen
	# Need to run npm compile in the ./mainnet directory for build/contracts to be created.
	cat ./build/contracts/ABIGenTest.json | jq '.abi' > ./build/ABIGenTest.json
	./abigen --abi ./build/ABIGenTest.json --pkg abigentest --out ./abigen_test.go 
	rm abigen
