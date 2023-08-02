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

function main() {
  // Enter your code here
  interface AdvancedArithmetic {
    divisorSum(no: number): number
  }

  //the class definition
  class Calculator implements AdvancedArithmetic {
    divisorSum(no: number): number {
      let sumDivisor = no
      for (let i = 1; i < no; i++) {
        if (no % i === 0) {
          sumDivisor += i
        }
      }
      return sumDivisor
    }
  }
  const no = +readLine()
  const myCalc = new Calculator()
  console.log('I implemented: AdvancedArithmetic')
  console.log(myCalc.divisorSum(no))
}
