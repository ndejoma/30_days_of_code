'use strict'

var _input = ''
var _index = 0
process.stdin.on('data', data => {
  _input += data
})
process.stdin.on('end', () => {
  _input = _input.split(new RegExp('[ \n]+'))
  main()
})
function read() {
  return _input[_index++]
}

/**** Ignore above this line. ****/

class Person {
  /**
   * Creates an instance of Person.
   * @param {string} firstName
   * @param {string} lastName
   * @param {number[]} identification
   * @memberof Person
   */
  constructor(firstName, lastName, identification) {
    this.firstName = firstName
    this.lastName = lastName
    this.idNumber = identification
  }

  printPerson() {
    console.log(
      'Name: ' +
        this.lastName +
        ', ' +
        this.firstName +
        '\nID: ' +
        this.idNumber
    )
  }
}

class Student extends Person {
  /*
   *   Class Constructor
   *
   *   @param {string} firstName - A string denoting the Person's first name.
   *   @param lastName - A string denoting the Person's last name.
   *   @param id - An integer denoting the Person's ID number.
   *   @param scores - An array of integers denoting the Person's test scores.
   */
  // Write your constructor
  /**
   * Creates an instance of Student.
   * @param {string} firstName
   * @param {string} lastName
   * @param {number} id
   * @param {number[]} scores
   * @memberof Student
   */
  constructor(firstName, lastName, id, scores) {
    super(firstName, lastName, id)
    this.scores = scores
  }

  /*
   *   Method Name: calculate
   *   @return A character denoting the grade.
   */
  // Write your method here
  calculate() {
    let totalScore = 0
    for (let score of this.scores) {
      totalScore += score
    }
    const averageScore = Math.round(totalScore / this.scores.length)
    let grade = ''
    if (averageScore >= 90 && averageScore <= 100) {
      grade = 'O'
    } else if (averageScore >= 80 && averageScore < 90) {
      grade = 'E'
    } else if (averageScore >= 70 && averageScore < 80) {
      grade = 'A'
    } else if (averageScore >= 55 && averageScore < 70) {
      grade = 'P'
    } else if (averageScore >= 40 && averageScore < 55) {
      grade = 'D'
    } else {
      grade = 'T'
    }
    return grade
  }
}

function main() {
  let firstName = read()
  let lastName = read()
  let id = +read()
  let numScores = +read()
  let testScores = new Array(numScores)

  for (var i = 0; i < numScores; i++) {
    testScores[i] = +read()
  }

  let s = new Student(firstName, lastName, id, testScores)
  s.printPerson()
  s.calculate()
  console.log('Grade: ' + s.calculate())
}
