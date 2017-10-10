// @flow

const fs = require('fs')
const path = require('path')
const {spawnSync} = require('child_process')

if (process.argv.length != 3) {
  console.log('USAGE: node abigen.js <path to json contracts>')
  process.exit(1)
}

const validContracts = [
  "BridgeToken",
  "Root",
  "SWCqueue",
  "SweetToken",
  "SweetTokenLogic",
  "TokenLogic",
  "UserDirectory"

  // "Treasury",
  // "Vault",
  // "VaultConfig",
  // "Wallet"
]

const validContractsBin = [
  "UserDirectory"
  // "Vault",
  // "Wallet"
]

const currentDir = process.cwd()
const abiDir = path.join(__dirname, 'tmp-contracts-abi')
const goDir = path.join(__dirname, '..', 'go-contracts')
let contractsDir = process.argv[2]

if (contractsDir[0] != '/')
  contractsDir = path.join(currentDir, contractsDir)

if (fs.existsSync(abiDir)) {
  console.log('removing "abiDir" content: ', abiDir)
  fs.readdirSync(abiDir).forEach(
    fname => fs.unlinkSync(path.join(abiDir, fname)))
} else {
  fs.mkdirSync(abiDir)
}

function extractAbi() {
  for (let fname of fs.readdirSync(contractsDir)) {
    if (!fname.endsWith('.json'))
      continue
    var cname = fname.slice(0, -5)
    if (validContracts.indexOf(cname) < 0 ) {
      // console.log("ignoring", fname)
      continue
    }

    var p = path.join(contractsDir, fname)
    var contract = require(p)
    console.log('extracting ABI', contract.contractName)

    fs.writeFileSync(path.join(abiDir, fname + '.abi'), JSON.stringify(contract.abi))
    if (validContractsBin.indexOf(cname) < 0)
      continue
    if (! contract.bytecode)
      console.error('Contract ', fname, ' doesn\'t have bytecode')
    else {
      console.log('extracting BIN', contract.contractName)
      fs.writeFileSync(path.join(abiDir, fname + '.bin'), JSON.stringify(contract.bytecode))
    }
  }
}

function generateGoInterfaces() {
  console.log('writing go bindings to ' + goDir)
  for (let fname of fs.readdirSync(abiDir)) {
    if (!fname.endsWith('.abi'))
      continue

    console.log('abigen', fname)
    var cname = fname.slice(0, -9)  // remove .json.abi
    var params = ['-abi', path.join(abiDir, fname),
                  '-type', cname,
                  '-out', path.join(goDir, cname+'.go'),
                  '-pkg', 'contracts']
    if (validContractsBin.indexOf(cname) >= 0)
      params = params.concat(['-bin', path.join(abiDir, fname.slice(0, -3) + 'bin')])
    var ret = spawnSync('abigen', params)
    if (ret.status != 0)
      console.log(ret.status, ret.stdout.toString(), ret.stderr.toString())
  }
}

extractAbi()
generateGoInterfaces()
