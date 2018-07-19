Example of how the cyclic chained queue works...

```
 bucket: [ _, _, _ ] --- linked bucket = nil
            |         
            |         
             ----------- popI = pushI = 0, full = false

# add an item (x)
  bucket: [ x, _, _ ] --- linked bucket = nil
            |  |       
            |  |       
            |   --------- pushI = 1
             ------------ popI = 0

# add an item (y)
  bucket: [ x, y, _ ] -- linked bucket = nil
            |     |    
            |     |    
            |      ------ pushI = 2
             ------------ popI = 0

# add more items (z, a, b)
  bucket: [ x, y, z ] --- [ a, b, _ ] - 
            |               |     |
            |               |      -- pushI = 2
            |                -------- popI = 0, full = false
            |
             ---------------- pushI = popI = 0, full = true

# pop an item
  bucket: [ _, y, z ] --- [ a, b, _ ] - 
            |  |            |     |
            |  |            |      -- pushI = 2
            |  |             -------- popI = 0
            |  |
            |   ------------- pushI = 0
             ---------------- popI = 1
               

# pop two more items
                        
  bucket: [ _, _, _ ] -X- [ a, b, _ ] - 
                  |    |    |     |
                  |    |    |      -- pushI = 2
                  |    |     -------- popI = 0
                  |    |
                   ----------- pushI = popI = 2, full = false
                       |
                        --- change start bucket pointer to next bucket
```
