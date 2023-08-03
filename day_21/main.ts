'use strict'

process.stdin.resume()
process.stdin.setEncoding('utf-8')
let inputString: string = ''
let inputLines: string[] = []
let currentLine: number = 0
process.stdin.on('data', function (inputStdin: string): void {
  inputString += inputStdin
})

process.stdin.on('end', function (): void {
  inputLines = inputString.split('\n')
  inputString = ''
  main()
})

function readLine(): string {
  return inputLines[currentLine++]
}
function printArray<Type>(arr: Type[]) {
  for (let el of arr) {
    console.log(el)
  }
}

function main() {
  // Enter your code here
  //the number of arrays to call
  let n = +readLine()
  for (let i = 0; i < n; i++) {
    const arr = readLine().split(' ').map(Number)
    printArray(arr)
  }
  n = +readLine()
  for (let i = 0; i < n; i++) {
    const arr = readLine().split(' ')
    printArray(arr)
  }
}
