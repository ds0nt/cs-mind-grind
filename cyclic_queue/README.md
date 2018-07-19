Example of how the cyclic chained queue works...

```
 bucket: [ _, _, _ ] -
            |         |
            |         -- linked bucket = nil
             ----------- popI = pushI = 0, full = false

# add an item (x)
  bucket: [ x, _, _ ] -
            |  |       |
            |  |       -- linked bucket = nil
            |   ---------- pushI = 1
             ------------ popI = 0

# add an item (y)
  bucket: [ x, y, _ ] -
            |     |    |
            |     |    -- linked bucket = nil
            |      ---------- pushI = 2
             ------------ popI = 0

# add more items (z, a, b)
  bucket: [ x, y, z ] - 
            |          |
            |          -- linked bucket = [ a, b, _ ] ...
             ---------------- pushI = popI = 0, full = true

# pop an item
  bucket: [ _, y, z ] -
            |  |       |
            |  |       -- linked bucket = [ a, b, _ ]
            |   ---------- popI = 1
             ------------ pushI = 0

# pop two more items
  bucket: [ a, b, _ ] - pop finals item; bucket = linked bucket
            |     |    |
            |     |    -- linked bucket = nil
            |      ---------- pushI = 2
             ------------ popI = 0
```
