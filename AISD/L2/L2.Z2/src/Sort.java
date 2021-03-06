import java.util.Random;

public class Sort {
	
	static int t[];
	static int insertionCompareKeys;
	static int insertionSwapKeys;
	static int mergeCompareKeys;
	static int mergeSwapKeys;
	static int quickCompareKeys;
	static int quickSwapKeys;
	
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
	
	public static int[] mergeSort(int t[], int n) {
		
		if(n == 0) {
			return t;
		}
		if (n == 1) {
			return t;
		} else {
			int half = n/2;
			int shalf = n-half;
			int[] A = new int[half];
			int[] B = new int[shalf];
			for(int i = 0; i < n; i++){
				if(i<half) {
					A[i]=t[i];
				} else {
					B[i-half] = t[i];
				}
			}
			t = merge(mergeSort(A,half), half, mergeSort(B,shalf),shalf);
		}
		return t;
	}
	
	public static int[] mergeSort2(int t[], int n) {
		
		if(n == 0) {
			return t;
		}
		if (n == 1) {
			return t;
		} else if(n < 10){
			insertionSort(t,0,n);
		} else {
			int half = n/2;
			int shalf = n-half;
			int[] A = new int[half];
			int[] B = new int[shalf];
			for(int i = 0; i < n; i++){
				if(i<half) {
					A[i]=t[i];
				} else {
					B[i-half] = t[i];
				}
			}
			t = merge(mergeSort(A,half), half, mergeSort(B,shalf),shalf);
		}
		return t;
	}
	
	public static int[] merge(int A[], int dlA, int B[], int dlB) {
		
		int[] ret = new int[dlA+dlB];
		int j = 0;			// indeks tablicy A
		int k = 0;			// indeks tablicy B
		while(j < dlA && k < dlB) {
			if(A[j]<B[k]) {
				mergeCompareKeys++;
				ret[j+k] = A[j];
				mergeSwapKeys++;
				j++;
			} else if (A[j] > B[k]) {
				mergeCompareKeys++;
				ret[j+k] = B[k];
				mergeSwapKeys++;
				k++;
			} else {
				ret[j+k] = A[j];
				mergeSwapKeys++;
				j++;
			}
		}
		
		// wiemy ze pierwsza tablica juz skonczyla a w drugiej sa same wieksze elementy
		
		while(k < dlB) {
			mergeSwapKeys++;
			ret[j+k] = B[k];
			k++;
		}
	
		while(j < dlA) {
			mergeSwapKeys++;
			ret[j+k] = A[j];
			j++;
		}
		
		return ret;
	}
	
	public static void quickSort(int[] A, int low, int high) {
		
		//partition
		//rekurencynie dala dwoch tablic
		
		int indexPivot = partition(A,low,high);
		if(indexPivot == low) return;
		quickSort(A,low,indexPivot-1);
		quickSort(A, indexPivot+1, high);
		
		
	}
	
	public static void drukuj(int[] t) {
		for(int i : t) {
			System.out.print(i+" ");
		}
		System.out.print("\n");
	}
	
	public static void quickSort2(int[] A, int low, int high) {
		
		if((high-low) < 5 ) {
			insertionSort(A,low,high);
			return;
		}
		//partition
		//rekurencynie dala dwoch tablic
		
		int indexPivot = partition(A,low,high);
		if(indexPivot == low) return;
		quickSort(A,low,indexPivot-1);
		quickSort(A, indexPivot+1, high);
		
		
	}
	
	public static int partition(int[] A, int p, int q) {
		if(p==q) return p;
		int pivot  = A[p];
		// i to indeks pivota
		int i = p, tmp;
		for(int j=p+1; j<q; j++) {
			quickCompareKeys++;
			if(A[j] <= pivot) {
				
				//swap A[j],a[i+1]
				quickSwapKeys++;
				i++;
				tmp = A[j];
				A[j] = A[i];
				A[i] = tmp;		
			}
		}
		quickSwapKeys++;
		tmp = A[p];
		A[p] = A[i];
		A[i] = tmp;
		
		return i;
	}
	
	public static int[] random(int N) {
		
		int ret[] = new int[N];
		
		Random r = new Random();
		
		for(int i = 0; i< ret.length; i++) {
			ret[i] = r.nextInt();
		}
		//drukuj(t);		
		return ret;
	}
	
	public static int[] desc(int N) {
		
		int ret[] = new int[N];
		
		Random r = new Random();
		
		if(ret.length == 0) {
			return ret;
		} else {
			ret[0] = r.nextInt();
			for(int i = 1; i < ret.length; i++){
				ret[i] = ret[i-1]-2;
			}
		}
		//drukuj(t);
		return ret;
	}
	
	
	public static void main(String[] args) {

		
		
		for(int i = 0; i < 4100; i+=100) {
			int comp = 0;
			int swap = 0;
			
			for(int j = 0; j<10; j++) {
				insertionCompareKeys = 0;
				insertionSwapKeys = 0;
//				quickCompareKeys = 0;
//				quickSwapKeys = 0;
				mergeCompareKeys = 0;
				mergeSwapKeys = 0;
				
				t = desc(i);
//				t = random(i);
				
				
//				quickSort2(t,0,t.length);
				mergeSort2(t,t.length);
//				comp += insertionCompareKeys+ quickCompareKeys;
//				swap += insertionSwapKeys + quickSwapKeys;
				comp += insertionCompareKeys+ mergeCompareKeys;
				swap += insertionSwapKeys + mergeSwapKeys;
			}
			comp /= 10;
			swap /= 10;
			System.out.print(i+" "+comp + " "+swap);
			comp = 0;
			swap = 0;
			for(int j = 0; j<100; j++) {
//				quickCompareKeys = 0;
//				quickSwapKeys = 0;
				mergeCompareKeys = 0;
				mergeSwapKeys = 0;
				
				t = desc(i);
//				t = random(i);
//				quickSort(t,0,t.length);
				mergeSort(t,t.length);
//				comp += quickCompareKeys;
//				swap += quickSwapKeys;
				comp += mergeCompareKeys;
				swap += mergeSwapKeys;
			}
			comp /= 10;
			swap /= 10;
			System.out.print(" "+comp + " "+swap);
			System.out.print("\n");
		}
		
		
	}
	

}
