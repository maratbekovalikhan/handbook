# Example Course Data

You can use the following data to test the "Add Course" functionality.

## Course Details

**Course Title:**
Mastering Go: From Zero to Hero

**Difficulty Level:**
Beginner to Intermediate

**Photo URL:**
https://go.dev/images/go-logo-white.svg

**Short Description:**
A comprehensive guide to learning the Go programming language, covering everything from basic syntax to advanced concurrency patterns.

**General Information:**
This course is designed for developers who want to build efficient, reliable, and scalable software. We will explore the philosophy of Go and build real-world applications.

---

## Course Sections (Curriculum)

### Section 1
**Title:**
Introduction to Go & Setup

**Information:**
Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.

In this section, we will:
1. Download and install Go from go.dev.
2. Setup your workspace and GOPATH.
3. Write your first "Hello, World!" program.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Section 2
**Title:**
Variables and Data Types

**Information:**
Go is a statically typed language. This means you must declare the type of variables, or let the compiler infer them.

Key concepts:
- `var` keyword
- Short variable declaration `:=`
- Basic types: `int`, `float64`, `bool`, `string`

Example:
```go
var name string = "Gopher"
age := 10
```

### Section 3
**Title:**
Control Structures (Loops & If/Else)

**Information:**
Go has only one looping construct: the `for` loop. It can be used as a while loop or a traditional C-style for loop.

We will also cover `if/else` statements and `switch` cases.

```go
for i := 0; i < 5; i++ {
    if i%2 == 0 {
        fmt.Println(i, "is even")
    }
}
```

### Section 4
**Title:**
Concurrency: Goroutines & Channels

**Information:**
One of Go's strongest features is its built-in support for concurrency.

- **Goroutines**: Lightweight threads managed by the Go runtime. Started with the `go` keyword.
- **Channels**: Typed conduits that allow goroutines to communicate with each other and synchronize their execution.

```go
func say(s string) {
    for i := 0; i < 5; i++ {
        fmt.Println(s)
    }
}

func main() {
    go say("world")
    say("hello")
}
```

---

# Example Course 2: React

## Course Details

**Course Title:**
React.js: The Complete Guide

**Difficulty Level:**
Intermediate

**Photo URL:**
https://upload.wikimedia.org/wikipedia/commons/thumb/a/a7/React-icon.svg/1200px-React-icon.svg.png

**Short Description:**
Learn React.js from scratch! Build modern, interactive web applications with the most popular JavaScript library.

**General Information:**
This course covers everything you need to know about React, including Hooks, Redux, React Router, and Next.js. Perfect for frontend developers looking to level up.

## Course Sections (Curriculum)

### Section 1
**Title:**
What is React? & JSX

**Information:**
React is a JavaScript library for building user interfaces. It uses a component-based architecture.

**JSX** is a syntax extension for JavaScript that looks like HTML. It makes writing React components intuitive.

```jsx
const element = <h1>Hello, world!</h1>;

function Welcome(props) {
  return <h1>Hello, {props.name}</h1>;
}
```

### Section 2
**Title:**
Components & Props

**Information:**
Components are the building blocks of any React application. They let you split the UI into independent, reusable pieces.

**Props** (short for properties) are how you pass data from a parent component to a child component.

```jsx
function App() {
  return (
    <div>
      <Welcome name="Sara" />
      <Welcome name="Cahal" />
      <Welcome name="Edite" />
    </div>
  );
}
```

### Section 3
**Title:**
State & Hooks (useState)

**Information:**
State allows React components to change their output over time in response to user actions, network responses, and anything else.

The `useState` Hook lets you add React state to function components.

```jsx
import React, { useState } from 'react';

function Counter() {
  const [count, setCount] = useState(0);

  return (
    <div>
      <p>You clicked {count} times</p>
      <button onClick={() => setCount(count + 1)}>
        Click me
      </button>
    </div>
  );
}
```

### Section 4
**Title:**
Effect Hook (useEffect)

**Information:**
The `useEffect` Hook lets you perform side effects in function components. It serves the same purpose as `componentDidMount`, `componentDidUpdate`, and `componentWillUnmount` in React classes.

Common uses: data fetching, subscriptions, or manually changing the DOM.

```jsx
import React, { useState, useEffect } from 'react';

function Example() {
  const [count, setCount] = useState(0);

  useEffect(() => {
    document.title = `You clicked ${count} times`;
  });

  return (
    <div>
      <p>You clicked {count} times</p>
      <button onClick={() => setCount(count + 1)}>
        Click me
      </button>
    </div>
  );
}
```
