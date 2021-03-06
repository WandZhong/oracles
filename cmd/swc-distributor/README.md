# swc-distributor

SWC distributor uses a provided credentials to the account holding enough SWT to distribute them according to the distribution list specified as an argument to the command. The distribution list file is a CSV file with same structure as `example.csv`.

The more specific documentation:

* [Release Note - SWC Distributor](https://drive.google.com/open?id=1XHzgqnl4VJ8AosKwRH9IR1wzLXMt9NjJ_JZG2NIdI00)
* [Use-case description](https://docs.google.com/document/d/1r1mE4FJhasBzjtwcx4ApQUXMe61FZVLzTwLfb5Uw8T0/edit#heading=h.jpodtnqw5wn1)

## Usage

Run command with `-h` to check the usage:

	./bin/swc-distributor -h

The program requires path to the contract schema directory (generated by truffle).

+ if the `-md5sum` argument is specified the program will check if the input file has correct control sum.
+ if a cell in the done column has some positive value (anything except: `""` (empty), `0`, `no`, false) then the whole row will be ignored
+ the application detects duplication. If there are two rows with the same address then it will return an error.


## TODO

+ [ ] crosscheck for correct address
+ [ ] KYC check
