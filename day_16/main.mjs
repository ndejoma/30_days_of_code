// 'use strict'

// process.stdin.resume()
// process.stdin.setEncoding('utf-8')

// let inputString = ''
// let currentLine = 0

// process.stdin.on('data', function (inputStdin) {
//   inputString += inputStdin
// })

// process.stdin.on('end', function () {
//   inputString = inputString.split('\n')

//   main()
// })

// function readLine() {
//   return inputString[currentLine++]
// }

// function main() {
//   const s = readLine()
//   try {
//     let error = new Error("Bad string")
//     let num = Number.parseInt(s) || error
//     console.log(num)
//   } catch (err) {
//     console.log(err.message)
//   }
// }

process.stdin.resume();
process.stdin.setEncoding('ascii');
var input_stdin = "";
var input_stdin_array = "";
var input_currentline = 0;
process.stdin.on('data', function (data) {
    input_stdin += data;
});
process.stdin.on('end', function () {
    input_stdin_array = input_stdin.split("\n");
    main();
});
function readLine() {
    return input_stdin_array[input_currentline++];
}

function main() {
    let s = readLine();
    try {
        let num = Number.parseInt(s) || error;
        console.log(num);
    }
    catch (error) {
        console.log("Bad String");
    }
}
