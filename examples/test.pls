module main;  // Similar to #ifndef main #define main ... #endif

using "strings"; // STD Strings Library

// Main function does not need to return anything in Pulse
fn main () {

  // Variable Declarations         // C Type Comparison
  var x: int = 45;                 // int
  var intPtr: int*;                // int*
  var green: []int = [0, 255, 0];  // int [3]
  var name = "Pulse";              // char[6];

  intPtr = &x;
  x = 45;

  std.printf("x value access through pointer: %d\n", *intPtr);

  var str = string::new("Pulse Is Awesome");
  var firstChar  = str.at(0); // P

  printf("Length: %d\n", s.length());
  str.clear();
  str.set("Hello World!");
  printf("str: %s\n", str.raw());

  // Dynamic Memory Allocation
  var BUFFER_SIZE = 4096;
  var buffer = calloc (BUFFER_SIZE);

  // Do stuff with Memory

  free(buffer);

}

// Defining Types
struct Person {
  name: char*;
  age: int;
}

// Method Overloads
impl Person::speak () {
  printf("Hello!, My name is %s and I am %d years of age.\n", self.name, self.age);
}

// All methods are mutatable and would be the same as passing p in as a pointer
fn (p Person) Birthday () {
  p.age += 1;
}

// These two functions are the same but using diferent calling styles
fn birthday (p Person*) {
  p->age += 1;
}
