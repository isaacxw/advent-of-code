const util = require("util");
const fs = require("fs");
const readFile = util.promisify(fs.readFile);

// Gets the input file based on the module using this function.
// @returns string[] of lines
async function getInput() {
  // const modulePath = module.parent.filename;
  const inputPath = 'input.txt';

  try {
    const file = await readFile(inputPath);

    return file.toString().split("\n");
  } catch (e) {
    console.error('No input file found. Please add one.');
    process.exit(1);
  }
}

exports.getInput = getInput;