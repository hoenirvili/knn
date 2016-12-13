# KNN algorithm proof of concept

```bash
$ go run main.go
[*] KNN algorithm
[*] Loading records

Sepal.Length	  Sepal.Width	  Species	  

5.3	  3.7	  Setosa	  
5.1	  3.8	  Setosa	  
7.2	  3	  Virginica	  
5.4	  3.4	  Setosa	  
5.1	  3.3	  Setosa	  
5.4	  3.9	  Setosa	  
7.4	  2.8	  Virginica	  
6.1	  2.8	  Versicolor	  
7.3	  2.9	  Virginica	  
6	  2.7	  Versicolor	  
5.8	  2.8	  Virginica	  
6.3	  2.3	  Versicolor	  
5.1	  2.5	  Versicolor	  
6.3	  2.5	  Versicolor	  
5.5	  2.4	  Versicolor	  

[i] Give the X point to look for
x = 5.2

y = 3.1

[*] Using k as 1

[*] [{name:Setosa count:1}]

[X = 5.100000 Y = 3.300000, Distance = 0.223607 Label = Setosa
 X = 5.400000 Y = 3.400000, Distance = 0.360555 Label = Setosa
 X = 5.300000 Y = 3.700000, Distance = 0.608276 Label = Setosa
 X = 5.100000 Y = 2.500000, Distance = 0.608276 Label = Versicolor
 X = 5.800000 Y = 2.800000, Distance = 0.670820 Label = Virginica
 X = 5.100000 Y = 3.800000, Distance = 0.707107 Label = Setosa
 X = 5.500000 Y = 2.400000, Distance = 0.761577 Label = Versicolor
 X = 5.400000 Y = 3.900000, Distance = 0.824621 Label = Setosa
 X = 6.000000 Y = 2.700000, Distance = 0.894427 Label = Versicolor
 X = 6.100000 Y = 2.800000, Distance = 0.948683 Label = Versicolor
 X = 6.300000 Y = 2.500000, Distance = 1.252996 Label = Versicolor
 X = 6.300000 Y = 2.300000, Distance = 1.360147 Label = Versicolor
 X = 7.200000 Y = 3.000000, Distance = 2.002498 Label = Virginica
 X = 7.300000 Y = 2.900000, Distance = 2.109502 Label = Virginica
 X = 7.400000 Y = 2.800000, Distance = 2.220360 Label = Virginica
]
[*] Result for X is [*] X = 5.200000, Y = 3.100000 Label = Setosa

[*] Using k as 3

[*] [{name:Setosa count:3}]

[*] Result for X is [*] X = 5.200000, Y = 3.100000 Label = Setosa

[*] Using k as 5

[*] [{name:Setosa count:3} {name:Versicolor count:1} {name:Virginica count:1}]

[*] Result for X is [*] X = 5.200000, Y = 3.100000 Label = Setosa

[*] Using k as 7

[*] [{name:Setosa count:4} {name:Versicolor count:2} {name:Virginica count:1}]

[*] Result for X is [*] X = 5.200000, Y = 3.100000 Label = Setosa

[*] Using k as 9

[*] [{name:Setosa count:5} {name:Versicolor count:3} {name:Virginica count:1}]

[*] Result for X is [*] X = 5.200000, Y = 3.100000 Label = Setosa

[*] Using k as 11

[*] [{name:Setosa count:5} {name:Versicolor count:5} {name:Virginica count:1}]

[*] Result for X is [*] X = 5.200000, Y = 3.100000 Label = Setosa

```
