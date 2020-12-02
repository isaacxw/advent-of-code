const { getInput } = require('../generator');

const part1 = (input) => {
    for (let i = 0; i < input.length; i++) {
        const a = input[i];
        const b = input.find(x => x === 2020 - a);

        if (b) {
            return a * b;
        }
    }
    return null;
}

const part2 = (input) => {
    for (let i = 0; i < input.length; i++) {
        for (let j = 1; j < input.length; j++) {
            for (let k = 2; k < input.length; k++) {
                if (input[i] + input[j] + input[k] === 2020) {
                    return input[i] * input[j] * input[k];
                }
            }
        }
    }
    return null;
} 

async function main() {
    const input = (await getInput()).map(Number);
    
    const res1 = part1(input);
    console.log(res1);

    const res2 = part2(input);
    console.log(res2);
}

main();