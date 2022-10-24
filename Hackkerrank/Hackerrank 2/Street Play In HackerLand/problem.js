'use strict';

const fs = require('fs');

process.stdin.resume();
process.stdin.setEncoding('utf-8');

let inputString = '';
let currentLine = 0;

process.stdin.on('data', function (inputStdin) {
    inputString += inputStdin;
});

process.stdin.on('end', function () {
    inputString = inputString.split('\n');

    main();
});

function readLine() {
    return inputString[currentLine++];
}



/*
 * Complete the 'getWhiteLightLength' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER m
 *  3. 2D_INTEGER_ARRAY lights
 */

function getWhiteLightLength(n, m, lights) {
    let max = 0;
    let min = lights[0][2];
    let exist = [];

    for (let i = 0; i < lights.length; i++) {
        if (max < lights[i][1]) {
            max = lights[i][1]
            if (min > lights[i][2] && min != max) {
                min = lights[i][2]
            }
        }

        if (!exist.includes(lights[i][0])) {
            exist.push(lights[i][0])
        }
    }
    console.log(exist)

    if (exist.length == 3) {
        console.log(min, max)
    }
}

function main() {
    const ws = fs.createWriteStream(process.env.OUTPUT_PATH);

    const n = parseInt(readLine().trim(), 10);

    const m = parseInt(readLine().trim(), 10);

    const lightsRows = parseInt(readLine().trim(), 10);

    const lightsColumns = parseInt(readLine().trim(), 10);

    let lights = Array(lightsRows);

    for (let i = 0; i < lightsRows; i++) {
        lights[i] = readLine().replace(/\s+$/g, '').split(' ').map(lightsTemp => parseInt(lightsTemp, 10));
    }

    const result = getWhiteLightLength(n, m, lights);

    ws.write(result + '\n');

    ws.end();
}
