# Sort
**Internal Sort**: Sort operation only happend in memory.

**External Sort**: Sort operation on Disk, and need to cooperate with memory.
## Time Complexity

### Quick Sort
* Best Case: $O(nlogn)$

$$\begin{equation}\begin{aligned}
T(n) &\le cn + 2T(n/2), c\, constant \\
&\le cn + 2(cn/2+2T(n/4)) \\
&\le 2cn + 4T(n/4)\\
\vdots \\
&\le cnlog_2n+nT(1) = O(nlogn)
\end{aligned}
\end{equation}$$

* Worst Case: $O(n^2)$

Worst Case happend on element arange in reverse order

$$
n + (n-1) + \cdots + 1 = \frac{n}{2}(n+1) = O(n^2)
$$

* Average Case: $O(nlogn)$

$$
T(n) = \underset{1 \le s\le n}{Ave}(T(s)+T(n-s)) + cn
$$

$$
\begin{equation}
    \begin{aligned}
    &\underset{1 \le s\le n}{Ave}(T(s)+T(n-s)) + cn \\
    &= \frac{1}{n}\sum_{s=1}^{n}(T(s)+T(n-s)) \\
    &= \frac{1}{n}(T(1)+T(n-1)+T(2)+T(n-2)+\cdots+T(n)+T(0))
    \end{aligned}
\end{equation}
$$

Because $T(0)=0$,

$$
T(n)=\frac{1}{n}(2T(1)+2T(2)+\cdots+2T(n-1)+T(n))+cn
$$

or,

$$
(n-1)T(n) = 2T(1)+2T(2)+\cdots+2T(n-1)+cn^2
$$

Therefore,

$$(n-1)T(n)-(n-2)T(n-1)=2T(n-1)+c(2n-1)$$
$$(n-1)T(n)-nT(n-1)=c(2n-1)$$
$$\frac{T(n)}{n}=\frac{T(n-1)}{n-1}+c\left (\frac{1}{n}+\frac{1}{n-1}\right )$$

We have,

$$
\begin{equation}
    \begin{aligned}
\frac{T(n)}{n}&=c\left(\frac{1}{n}+\frac{1}{n-1}+\cdot+\frac{1}{2}\right)+c\left(\frac{1}{n-1}+\frac{1}{n-2}+\cdot+1\right)\\
&=c(H_n-1)+c(H_{n-1})\\
&=c(H_n+H_{n-1}-1)\\
&=c(2H_n-\frac{1}{n}-1)\\
&=c(2H_n-\frac{n+1}{n})
    \end{aligned}
\end{equation}
$$

Finally, we have

$$
\begin{equation}
    \begin{aligned}
T(n)&=2cnH_n-c(n+1) \\
&\cong 2cnlog_en-c(n+1)\\
&=O(nlogn) 
    \end{aligned}
\end{equation}
$$

### Heap Sort
* Best Case: $O(nlogn)$
* Average Case: $O(nlogn)$
* Worst Case: $O(nlogn)$

Heap Sort has two phase:
1. Construct a max-heap(min-heap)
2. Swap Top element and Last element in heap and restore heap from Top element.

#### Time complexity for construct a max-heap:

$$ \sum_{L=0}^{d-1}2(d-L)2^L=2d\sum_{L=0}^{d-1}2^L-4d\sum_{L=0}^{d-1}L2^{L-1}$$

Let there $n$ be n numbers to be sorted, depth $d$ is $\lfloor logn \rfloor$. 

Let the level of an internal node be $L$, the maximum number of node at level $L$ is $2^L$

We will use this formula to prove:

$$
\begin{equation}
\sum_{L=0}^{k}L2^{L-1}=2^k(k-1)+1
\end{equation}
$$

Therefore,

$$
\begin{equation}
\begin{aligned}
\sum_{L=0}^{d-1}2(d-L)2^L &= 2d\sum_{L=0}^{d-1}2^L-4\sum_{L=0}^{d-1}L2^{L-1} \\
&=2d(2^d-1)-4(2^{d-1}(d-1-1)+1)\\
&=2d(2^d-1)-4(d2^{d-1}-2^d+1)\\
&=4 \times 2^d -2d-4 \\
&=4 \times 2^{\lfloor logn \rfloor}-\lfloor 2logn \rfloor-4\\
&=cn-\lfloor 2logn \rfloor-4\,\,\,\,where\,\,\,\, 2 \le c \le 4 \\
&\le cn = O(n)
\end{aligned}
\end{equation}
$$

#### Time complexity of delete Top element and restore heap
The total number comparison of delete Top element and restore heap from Top element:

$$
2\sum_{i=1}^{n-1}\lfloor logi \rfloor
$$

To evaluate this formula, let us consider the case of $n=10$

$$
\begin{equation}
\begin{aligned}
&\lfloor log1 \rfloor = 0 \\
&\lfloor log2 \rfloor = \lfloor log3\rfloor =1 \\
&\lfloor log4 \rfloor = \lfloor log5\rfloor =\lfloor log6 \rfloor = \lfloor log7\rfloor= 2 \\
&\lfloor log8 \rfloor = \lfloor log9\rfloor =3
\end{aligned}
\end{equation}
$$

We observe that herer are 

$2^1$ numbers equal to $\lfloor log2^1 \rfloor =1$

$2^2$ numbers equal to $\lfloor log2^2 \rfloor =2$

and $10 - 2^{\lfloor log10 \rfloor} = 10-2^3=2$ numbers equal to $\lfloor logn \rfloor$

In general,

$$
\begin{equation}
\begin{aligned}
2\sum_{i=1}^{n-1}\lfloor logi \rfloor &= 2\sum_{i=1}^{\lfloor logn \rfloor-1}i2^i+2(n-2^{\lfloor logn \rfloor})\lfloor logn \rfloor \\
&=4\sum_{i=1}^{\lfloor logn \rfloor -1}i2^{i-1}+2(n-2^{\lfloor logn \rfloor})\lfloor logn \rfloor
\end{aligned}
\end{equation}
$$

Using this formula in proving:

$$
\begin{equation}
\sum_{L=0}^{k}L2^{L-1}=2^k(k-1)+1
\end{equation}
$$

we have

$$
\begin{equation}
\begin{aligned}
&2\sum_{i=1}^{n-1} \lfloor logi \rfloor \\
&=4\sum_{i=1}^{\lfloor logn \rfloor-1}i2^{i-1}+2(n-2^{\lfloor logn \rfloor})\lfloor logn \rfloor \\
&=4(2^{\lfloor logn \rfloor-1}(\lfloor logn \rfloor-1-1)+1)+2n\lfloor logn \rfloor-2\lfloor logn \rfloor 2^{\lfloor logn \rfloor} \\
&=2 \cdot 2^{\lfloor logn \rfloor}\lfloor logn \rfloor-8 \cdot 2^{\lfloor logn \rfloor-1} +4+2n\lfloor logn \rfloor-2\cdot2^{\lfloor logn \rfloor}\lfloor logn \rfloor \\
&=2 \cdot n\lfloor logn \rfloor-4 \cdot 2^{\lfloor logn \rfloor}+4 \\
&= 2n\lfloor logn \rfloor - 4cn + 4 \,\,\, where \,\,\, 2 \le c \le 4 \\
&=O(nlogn)
\end{aligned}
\end{equation}
$$


