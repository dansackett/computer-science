# Factory Design Pattern

The factory design pattern is a way to decouple code through the use of a
common interface and similiar objects. It typically consists of:

- A factory function that can generate sub-classes
- A common interface which sub-classes conform to

With these we can create objects without exposing the creation logic to the
client and refer to newly created objects using the common interface.
