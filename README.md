### Data Transformar

Given 

•  input -  array of json objects

•  input - pivot columns

•  input - summarization columns and functions

•  output - a pivot table


```
               Jan              Feb            March
          Revenue Tax      Revenue Tax     Revenue Tax
 Laptop . 40000   100      10000   400     20000   200
 Ipad     30000   200      8000    50      30000   350
  {
     category: {
        laptop: {
            month: {
                jan: {
                    revenue: 40000,
                    tax: 100
                },
                feb: {
                    :
                }
            }
        },
        ipad: {
            month: {
                jan: {
                    revenue: 30000
                    tax: 350
                },
                feb: {
                }
            }
        }
     }
 }
 ```
 
 If more than one dimension/groupBy given
 e.g.
  dimension= ['category', 'rating']
  groupBy = ['year', 'month']
  measures = [{name: 'revenue', fun: 'sum'},{name: 'tax', fun: 'sum'}]
  Then output should be like:

  ```
                 Jan19               Feb19       ....          Jan20
            Revenue  Tax        Revenue  Tax              Revenue  Tax
 Laptop-1   40000    500        10000    100     ....    40000     500
 Laptop-2   40000    400        10000    200     ....    30000     600
 Laptop-3   40000    300        10000    300     ....    20000     600
  :
  :
  :
 Ipad-1     30000    200        4000    500      ....    13000     700
 Ipad-2     20000    300        5000    400      ....    33000     600
 Ipad-3     60000    100        6000    300      ....    12300     450
 Output JSON:
    {
        category : {
            laptop: {
                rating: {
                    1: {
                        year: {
                            2019: {
                                month: {
                                    jan: {
                                        revenue: 40000,
                                        tax: 300
                                    },
                                    feb: {
                                        revenue: 50000,
                                        tax: 400
                                    }
                                }
                            }
                        }
                    },
                    2: {
                    }
                }
            },
            ipad: {
            }
        }
    }
}
```
