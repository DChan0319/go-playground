package structs_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
  rectangle := Rectangle{10.0, 10.0}
  var actual float64 = rectangle.Perimeter()
  var expected float64 = 40.0

  if actual != expected {
    t.Errorf("actual: %g expected: %g", actual, expected)
  }
}

func TestArea(t *testing.T) {
  areaTests := []struct {
    name     string
    shape    Shape
    expected float64
  } {
    {name: "Rectangle", shape: Rectangle{Width: 12.0, Height: 6.0}, expected: 72.0},
    {name: "Circle", shape: Circle{Radius: 5.0}, expected: 78.53981633974483},
    {name: "Triangle", shape: Triangle{Height: 5.0, Base: 10.0}, expected: 25.0},
  }

  for _, tC := range areaTests {
    t.Run(tC.name, func(t *testing.T) {
      actual := tC.shape.Area()

      if actual != tC.expected {
        t.Errorf("struct: %+v, actual: %g expected: %g", tC.shape, actual, tC.expected)
      }
    })
  }
}
