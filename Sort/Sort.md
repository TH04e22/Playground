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