//
// Created by yangzhengkun on 2020/7/24.
//

#include <queue>
#include <string>
#include <utility>
#include <atomic>

class Pet {
 private:
  std::string type;

 public:
  explicit Pet(std::string t) : type(std::move(t)) {}
  std::string getType() {
	return type;
  }
};

class Cat : public Pet {
 public:
  explicit Cat(std::string t) : Pet(std::move(t)) {}
};

class Dog : public Pet {
 public:
  explicit Dog(std::string t) : Pet(std::move(t)) {}
};

class PetItem {
 private:
  Pet pet;
  int num;//编号
 public:
  PetItem(Pet p, int n) : pet(p), num(n) {}
  Pet getPet() { return pet; }
  int getNum() { return num; }
  std::string getType() { return pet.getType(); }
};

class CatDogQueue {
 private:
  std::queue<PetItem> dogQueue;
  std::queue<PetItem> catQueue;
  std::atomic<int> count;

 public:
  CatDogQueue() = default;
  void add(Pet p) {
	if (p.getType() == "dog") {
	  dogQueue.emplace(p, ++count);
	} else if (p.getType() == "cat") {
	  catQueue.emplace(p, ++count);
	}
  }

  Pet top() {
	if (!dogQueue.empty() && !catQueue.empty()) {
	  if (dogQueue.front().getNum() < catQueue.front().getNum()) {
		return dogQueue.front().getPet();
	  } else {
		return catQueue.front().getPet();
	  }
	} else if (!dogQueue.empty()) {
	  return dogQueue.front().getPet();
	} else {
	  return catQueue.front().getPet();
	}
  }

  Dog frontDog() {
	if (!dogQueue.empty()) {
	  return (Dog)dogQueue.front().getPet();
	}
  }
};