# Memory Access Sychronization

### Deadlock
`deadlock`: "program in which all the concurrent processes are waiting for one another" [Example](./deadlock/main.go)

Coffman Conditions (1971 paper that outlines conditions that must be present for a deadlock to arrise)
- Mututal Exclusion
> "A concurrent process holds exclusive rights toa resource at any time"
- Wait For Condition
> "A concurrent process must simultaneously hold a resource and be waiting for an additional resource"
- No Preemption
> "A resource held by a concurrent process can only be released by that process, so it fulfills this condition"
- Circular Wait
> "A concurrent process (P1) must be waiting on a chain of other concurrent processes (P2), which are in turn waiting on it (P1), so it fulfills this final condition too."


### Livelock
`livelock`: "Program that are actively performing concurrent operations, but these operations do nothing to move the state of the program forward" [Example](./livelock/main.go)
> livelocks are a subset of a larger problem set; starvation. IN the example, each resource was starved of a shared lock.


### Starvation
`starvation`: "any situation where a concurrent process cannot get all the resources it needs to perform work."
> a 'metric' is used to help detect starvation. i.e. logging when work is accomplished then determining if the rate of work is as high as expected.  Balance: coarse-grained (greedy worker) and fine-grained (polite worker) -- the performance of the greedy worker is higher but at the cost of the polite worker not being able to accomplish as much.