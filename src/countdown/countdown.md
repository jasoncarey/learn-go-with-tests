If your mocking code is becoming complicated or you are having to mock out lots of things to test something, you should listen to that bad feeling and think about your code. Usually it is a sign of:

- The thing you are testing is having to do too many things (because it has too many dependencies to mock)

    - Break the module apart so it does less

- Its dependencies are too fine-grained

    - Think about how you can consolidate some of these dependencies into one meaningful module

- Your test is too concerned with implementation details

    - Favour testing expected behaviour rather than the implementation

Normally a lot of mocking points to bad abstraction in your code.

Don't test private functions

When faced with less trivial examples, break the problem down into "thin vertical slices". Try to get to a point where you have working software backed by tests as soon as you can.