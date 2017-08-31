// @flow

const fs = require('fs')
const path = require('path')
const {spawnSync} = require('child_process')

if (process.argv.length != 3) {
  console.log('USAGE: node abigen.js <path to json contracts>')
  process.exit(1)
}

const currentDir = process.cwd()
const abiDir = path.join(__dirname, 'tmp-contracts-abi')
const goDir = path.join(__dirname, '..', 'go-contracts')
let contractsDir = process.argv[2]
if (contractsDir[0] != '/')
  contractsDir = path.join(currentDir, contractsDir)

if (!fs.existsSync(abiDir)){
  fs.mkdirSync(abiDir)
}


function extractAbi() {
  for (let fname of fs.readdirSync(contractsDir)) {
    if (!fname.endsWith('.json'))
      continue
    var p = path.join(contractsDir, fname)
    var contract = require(p)
    console.log("extracting ABI and BIN", contract.contract_name)

    fs.writeFileSync(path.join(abiDir, fname + '.abi'), JSON.stringify(contract.abi))
    fs.writeFileSync(path.join(abiDir, fname + '.bin'), JSON.stringify(contract.binary || contract.unlinked_binary))
  }
}

function generateGoInterfaces() {
  for (let fname of fs.readdirSync(abiDir)) {
    if (!fname.endsWith('.abi'))
      continue
    if (fname.indexOf('Lib.json') >= 0 || fname.indexOf('Events.json') ) {
      console.log("ignoring", fname)
      continue
    }

    console.log('abigen', fname)
    var name = fname.slice(0, -9)  // remove .json.abi
    ret = spawnSync('abigen', ['-abi', path.join(abiDir, fname),
                               '-bin', path.join(abiDir, fname.slice(0, -3) + 'bin'),
                               '-type', name,
                               '-out', path.join(goDir, name+'.go'),
                               '-pkg', 'contracts'])
    if (ret.status != 0)
      console.log(ret.status, ret.stdout.toString(), ret.stderr.toString())
  }
}

extractAbi()
generateGoInterfaces()
