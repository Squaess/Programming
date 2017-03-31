
public class DualPivot {
	static int insertionCompareKeys;
	static int insertionSwapKeys;
	
	public static void insertionSort(int t[], int low, int high) {
	
		
		if(low-high == 0) {
			return;
		
		}
		for (int i = low+1; i < high; i++) {
			int key = t[i];
			int j = i-1;
			while(j >= low && key < t[j]) {
				t[j+1]=t[j];
				j--;
				insertionCompareKeys++;
				insertionSwapKeys++;
			}
			
			t[j+1] = key;
		}
	}

	public static void QDP(int[] A, int low, int high) {
		int length = high-low;
		if(length < 17) {
			insertionSort(A,low,high);
			return;
		} else {
			
		}
	}
}
