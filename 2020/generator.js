const util = require("util");
const fs = require("fs");
const readFile = util.promisify(fs.readFile);

// Gets the input file based on the module using this function.
async function getInput() {
  const inputPath = 'input.txt';

  try {
    const file = await readFile(inputPath);
    
    return lines = file.toString().split("\n");
  } catch (err) {
    console.error('No input file found.');
    process.exit(1);
  }
}

exports.getInput = getInput;