/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
   L, R := 1, n

   var mid int

   for L <= R {
		mid = (L + R) / 2

		if guess(mid) == -1 {
			R = mid - 1
		} else if guess(mid) == 1 {
			L = mid + 1
		} else {
			return mid
		}
   }

   return -1
}
