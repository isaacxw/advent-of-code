const { getInput } = require('../generator');

/**
 * Parse file into workable format
 * 
 * @param {string} line 
 */
const parseInput = (line) => {
    const [policy, password] = line.split(": ");
    const [range, letter] = policy.split(" ");

    return {
        password,
        range: range.split("-").map(Number),
        letter,
    };
};

/**
 * Count the number of occurences the letter appears
 * in the string
 * 
 * @param {*} policy
 * @return {boolean}
 */
const policyRange = policy => {
    let [min, max] = policy.range;
    let count = 0;

    for (let i = 0; i < policy.password.length; i++) {
        if (policy.password.charAt(i) === policy.letter) {
            count++;
        }
    }

    return count >= min && count <= max;

}

/**
 * Part 1 result
 * 
 * @returns {number}
 */
const part1 = policies => {
    return policies.filter(policyRange).length;
}

/**
 * Determine if policy holds for part 2
 * 
 * @param {*} policy
 * @returns {boolean} 
 */
const passwordPosReq = policy => {
    const [ pos1, pos2 ] = policy.range;

    const passwordHasLetterAt = index => policy.password[index - 1] === policy.letter;

    const atPos1 = passwordHasLetterAt(pos1);
    const atPos2 = passwordHasLetterAt(pos2);
    return (atPos1 || atPos2) && !(atPos1 && atPos2);
}

/**
 * Part 2 solution
 * 
 * @param {[]} policies 
 * @returns {number}
 */
const part2 = policies => {
    return policies.filter(passwordPosReq).length;
}

const main = async () => {
    const input = (await getInput()).map(parseInput);

    // 422
    const res1 = part1(input);
    console.log("Part 1: " + res1);

    // 451
    const res2 = part2(input);
    console.log("Part 2: " + res2);
}

main();
