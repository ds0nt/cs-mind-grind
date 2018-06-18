package main

func main() {
	println("https://www.geeksforgeeks.org/hashing-set-1-introduction/")
	println("https://www.geeksforgeeks.org/hashing-set-2-separate-chaining/")
	println("https://www.geeksforgeeks.org/birthday-paradox/")
	println("taylor series... https://en.wikipedia.org/wiki/Taylor_series -- wow, I forget how to math things.. calculus + probability + infinite sums.. all at once yikes")
	println("consider basically refreshing your brain on calc + statistics. which you should have payed attention to in school")
	println("I will come back to this later.. back to hash-tables..")
	println("separate chaining is straightforward. colliding hashed keys go into linked lists.")
	println("seems pretty bad in reality, due to cache locality problems")
	println("https://www.geeksforgeeks.org/hashing-set-3-open-addressing/ video helped - but lacking explanation of how to increase size of hash table and how operations still would work.. and how the hash functions (\% table size) change")
	println("should enlighten me: https://web.archive.org/web/19990903133921/http://www.concentric.net/~Ttwang/tech/primehash.htm")
	println("super interesting: https://golang.org/src/runtime/hashmap.go")
	println("I wonder if I can define key generation function makes the go hashmap implementation suck horribly. It seems to double its size, and not use prime table sizes")
}
