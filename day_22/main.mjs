// Start of function Node
function Node(data) {
  this.data = data
  this.left = null
  this.right = null
} // End of function Node

// Start of function BinarySearchTree
function BinarySearchTree() {
  this.insert = function (root, data) {
    if (root === null) {
      this.root = new Node(data)

      return this.root
    }

    if (data <= root.data) {
      if (root.left) {
        this.insert(root.left, data)
      } else {
        root.left = new Node(data)
      }
    } else {
      if (root.right) {
        this.insert(root.right, data)
      } else {
        root.right = new Node(data)
      }
    }

    return this.root
  }

  // Start of function getHeight
  this.getHeight = function (root) {
    // Add your code here
    let height = 1
    let height_left = -1
    let height_right = -1
    if (root.left) {
      height_left = this.getHeight(root.left)
    }

    if (root.right) {
      height_right = this.getHeight(root.right)
    }

    height += Math.max(height_left, height_right)
    return height
  } // End of function getHeight
} // End of function BinarySearchTree

process.stdin.resume()
process.stdin.setEncoding('ascii')

var _input = ''

process.stdin.on('data', function (data) {
  _input += data
})

process.stdin.on('end', function () {
  let tree = new BinarySearchTree()
  let root = null

  let values = _input.split('\n').map(Number)

  for (let i = 1; i < values.length; i++) {
    root = tree.insert(root, values[i])
  }

  console.log(tree.getHeight(root));
//   console.log(tree)
})

//to run this code
// echo "7\n3\n5\n2\n1\n4\n6\n7" | node main.mjs