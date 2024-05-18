[Home](/)

# Notes from A Philosophy of Software Design
* Read in 2024

My key takeaways from reading "A Philosophy of Software Design" by John Ousterhout. Those are mostly quotes from the book. 

#### The Nature of Complexity
- Complexity is anything related to the structure of a software system that makes it hard to understand and modify

#### Symptoms of Complexity
- 1) Change amplification: a seemingly simple change requires code modification in many different places
- 2) Cognitive load: the amount a developer needs to know in order to complete a task
- 3) Unknown unknowns: it is not obvious which piece of code should be modified to complete a task

#### Causes of Complexity
- Complexity is caused by dependencies and obscurity
- A dependency exist when a piece of code cannot be understood in isolation
- Obscurity occurs when information is not obvious

#### Working Code isn't Enough
- Tactical Programming: only try to make the code work. Creates design issues in the long term
- Strategic Programming: primary goal is to produce a great design

#### Modular Design
- The goal of modular design is to minimize dependencies between modules
- Think about modules in two parts: an interface and an implementation
- The interface describes "what" the module does, and the implementation describes "how"
- An abstraction is a simplified view of an entity, which omits unimportant details
- Interfaces should be designed to make the common case as simple as possible

#### Module should be deep
- Module depth is the ratio between how simple an interface is and how much value the implementation provides
- "Shallow modules" are modules where the interface is almost as complex as the interface: the abstraction provides very little value
- "Deep modules" have a interface that is very simple and hide a lot of complexity in the implementation
- A good question to ask is: "what is the simplest interface that covers all my needs ?"

#### Information Leakage
- Information leakage occurs when a design decision is reflected in multiple modules
- It is one of the most important red flag in software design

#### Different Layer, Different Abstraction
- Each layer should provide a different abstraction from the one above and below it

#### General vs Specialized modules
- More general purpose modules are usually deeper
- Mixing general purpose code with specialized for a particular use code is a red flag
- Each method should do one thing and do it completely: avoid conjoined methods (that need to be always called together)

#### Define errors out of existence
- Classes with a lot of exception have complexe interfaces, making them shallower
- The best way to eliminate exceptions handling is to design the API so that there are no exception to handle: exceptions are handled inside the module

#### Design it twice
- To increase the likelihood of designing a better system: design it twice

#### Good code should be self-documenting
- If users must read the code of a method in order to use it, then there is no abstraction
- Comments should describe things that are not obvious from the code

#### Consistency
- Consistency is a good way to reduce the overall complexity of a system by making the code more obvious
- Resist the urge to "improve" existing conventions. Having a better idea is not a sufficient excuse to introduce inconsistencies

#### Code Should be Obvious
- Software should be designed for ease of reading, not ease of writing

#### Designing for Performance
- Programmers' intuition about performance is unreliable, even for experience developers
- Making changes based on intuition is likely to waste time, make the softwake more complicated without improving performance
- Before making a change, measure the system's existing behavior
- If code has never been run, it doesn't work