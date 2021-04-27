### Building Nakama from source

1. update to the last go version
2. using sqlite as a database backend

e.g: run dev. build of Nakama with ```nk.db``` sqlite database:
    ```./src/nakama/build/dev/nakama -database.address nk.db```

Prerequisite:
  1. Installing ```jq``` for tools
  2. Installing ```nodejs```
  3. Installing ```dep``` (set up *GOBIN*)
  4. _protoc_ v3.3.0 (newer versions might not work)
