// class Animal {
//   constructor() {
//     if (new.target == Animal) {
//       throw new Error(
//         "Abstract classes can't be instantiated. They are only meant to be used as the bases for other classes"
//       )
//     }
//   }
//   say() {
//     throw new Error("Method 'say()' must be implemented.")
//   }
//   eat() {
//     console.log('eating')
//   }
// }
// const an = new Animal()
// class Cat extends Animal {
//   constructor(name, breed, age) {
//     super()
//     this.name = name
//     this.breed = breed
//     this.age = age
//   }
//   say() {
//     console.log('Meooooow!!!')
//   }
// }

// const nigelCat = new Cat('Bruno', 'Calico', 5)
// nigelCat.say()

//day 14: Abstract classes
// 'use strict'

// var _input = ''
// var _index = 0
// process.stdin.on('data', data => {
//   _input += data
// })
// process.stdin.on('end', () => {
//   _input = _input.split(new RegExp('\n'))
//   main()
// })
// function readLine() {
//   return _input[_index++]
// }

// /**** Ignore above this line. ****/

// class Book {
//   constructor(title, author) {
//     if (this.constructor === Book) {
//       throw new TypeError(
//         'Do not attempt to directly instantiate an abstract class.'
//       )
//     } else {
//       this.title = title
//       this.author = author
//     }
//   }

//   display() {
//     console.log("Implement the 'display' method!")
//   }
// }

// // Declare your class here.
// class MyBook extends Book {
//   /**
//    *   Class Constructor
//    *
//    *   @param {string} title The book's title.
//    *   @param {string} author The book's author.
//    *   @param {number} price The book's price.
//    *   @memberof MyBook
//    **/
//   // Write your constructor here
//   constructor(title, author, price) {
//     super(title, author)
//     this.price = price
//   }
//   /**
//    *   Method Name: display
//    *
//    *   Print the title, author, and price in the specified format.
//    **/
//   // Write your method here
//   display() {
//     console.log(`Title: ${this.title}`)
//     console.log(`Author: ${this.author}`)
//     console.log(`Price: ${this.price}`)
//   }
// }

// // End class

// function main() {
//   let title = readLine()
//   let author = readLine()
//   let price = +readLine()

//   let book = new MyBook(title, author, price)
//   book.display()
// }

//day 15: scope
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

  class Difference {
    protected elements: number[]
    public maximumDifference: number = -1000
    constructor(elements: number[]) {
      this.elements = elements
    }
    public computeDifference(): number {
      let maximumDifference = -1000
      for (let i = 0; i < this.elements.length; i++) {
        for (let j = i + 1; j <= this.elements.length; j++) {
          let arr: number[] = this.elements
          let tempMaxDifference = Math.abs(arr[i] - arr[j])
          if (tempMaxDifference > maximumDifference) {
            maximumDifference = tempMaxDifference
            this.maximumDifference = maximumDifference
          }
        }
      }
      return this.maximumDifference
    }
  }

  const n = +readLine()
  const elements = readLine().split(' ').map(Number)
  console.log(elements)
  const difference = new Difference(elements)
  console.log(difference.computeDifference())
}
