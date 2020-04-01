# AVL Tree

## Definition
An AVL tree is a binary tree that
either is empty or consists of two AVL subtrees, TL, and TR, whose heights differ by no more than 1, as shown below.
```go
|HL-HR| <= 1
```
An AVL Tree is a height-balanced binary search tree.

## AVL Balance factor
The height of the left subtree minus the height of the right subtree.
The balance factor for AVL Tree must be +1, -1 or 0.

## Balancing Trees
AVL Trees are balanced by rotating nodes.<br/>
```go
switch {
case Left of Left:
	Rotate Right
case Right of Right:
	Rotate Left
case Right of Left:
	Rotate Left
	Rotate Right
case Left of Right:
	Rotate Right
	Rotate Left
}
```


Reference:<br/>
[1]Gilberg Behrouz, A. Forouzan. "Data Structures: A Pseudocode Approach with C, Second Edition Richard F."