## BTree rules:
1. Root: 1 to 2t-1 keys => 0-2t
2. Non-root: t-1 to 2t-1 keys
3. Non-leaves: one more children than keys
4. Leaves: all at same level
5. We consider top-down version

## Algorithm:
### Insert:
1. Check root node, if it is not full, move to next node if exist
2. Check the node, if it is full or not:
  Yes:
  1. Split the node and put new elemenet.
  No:
  1. Move to next node.


http://www.di.ufpb.br/lucidio/Btrees.pdf

https://www.youtube.com/watch?v=I22wEC1tTGo
https://www.youtube.com/watch?v=s3bCdZGrgpA
