function Person(initialAge) {
  // Add some more code to run some checks on initialAge
  if (initialAge < 0) {
    console.log('Age is not valid, setting age to 0.')
    this.age = 0
  } else {
    this.age = initialAge
  }
}

Person.prototype.amIOld = function () {
  if (this.age < 13) {
    console.log('You are young.')
  } else if (this.age >= 13 && this.age < 18) {
    console.log('You are a teenager.')
  } else {
    console.log('You are old.')
  }
}

Person.prototype.yearPasses = function () {
  this.age += 1
}

const person1 = new Person(20)

console.log(`The age of person1 4 years ago was ${person1.age}`)
for (let i = 1; i <= 4; i++) {
  person1.yearPasses()
}
console.log(`The age of person 1 after 4 years is ${person1.age}`)
console.log(
  `Person 1 is asking: Am I old? and the answer is?`
)
person1.amIOld()