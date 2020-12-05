const { getInput } = require("../generator");

const part1 = input => {
    const area = input.map((line) => line.trim().split(""));

    let x = 0;
    let y = 0;
    let treeCount = 0; 

    while (y < area.length) {
        let x_range = x % area[y].length;
        if (area[y][x_range] === "#") {
            treeCount++;
        }

        x += 3;
        y += 1;
    }

    return treeCount;
}

const part2 = input => {
    const area = input.map((line) => line.trim().split(""));
    let total = 0;

    const findTrees = (moveX, moveY) => {
        let x = 0;
        let y = 0;
        let treeCount = 0;

        while (y < area.length) {
            let x_range = x % area[y].length;
            if (area[y][x_range] === "#") {
                treeCount++;
            }

            x += moveX;
            y += moveY;
        }

        return treeCount;
    }

    return findTrees(1, 1) * findTrees(3, 1) * findTrees(5, 1) * findTrees(7, 1) * findTrees(1, 2);
}

const main = async () => {
    const input = (await getInput()).map(String);

    // 252
    const res1 = part1(input);
    console.log(res1);

    // 2608962048
    const res2 = part2(input);
    console.log(res2);
}

main();