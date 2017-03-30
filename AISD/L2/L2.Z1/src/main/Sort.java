package main;

import java.util.Random;
import java.util.Scanner;

public class Sort {
	static int t[];
	static int insertionCompareKeys;
	static int insertionSwapKeys;
	static int mergeCompareKeys;
	static int mergeSwapKeys;
	static int quickCompareKeys;
	static int quickSwapKeys;
	
	/*
	 * Tablica danych t , n to romiar tablicy
	 */
	public static void insertionSort(int t[], int n) {
		insertionCompareKeys = 0;
		insertionSwapKeys = 0;
		
		if(n == 0) {
			return;
		
		}
		for (int i = 1; i < n; i++) {
			int key = t[i];
			int j = i-1;
			while(j >= 0 && key < t[j]) {
				t[j+1]=t[j];
				j--;
				insertionCompareKeys++;
				insertionSwapKeys++;
			}
			if(j >= 0) insertionCompareKeys++;
			insertionSwapKeys++;
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
	
	/***********************************************************************************
	 * napisac drugiego merga na kolejce
	 */
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
		//rekurencynie dala tw�ch tablic
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
	
	public static void drukuj(int[] t) {
		for(int i : t) {
			System.out.print(i+" ");
		}
		System.out.print("\n");
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
		
		
		int decision = 0;					//Wybor czy losowy ciag, czy malejacy
		int N = 0;							//rozmiar danych
		boolean warunek = true;
		
		Scanner input = new Scanner(System.in);
			
		while(warunek) {
			System.out.println("Wybierz odpowiednia liczbe: ");
			System.out.println("1. Losowy ciag");
			System.out.println("2. Ciag malejacy");
			decision = input.nextInt();
			switch(decision) {
			case 1:
				warunek = false;
				System.out.println("Podaj wielkosc danych: ");
				N = input.nextInt();
				test(N,true);
				// TODO zrob co nalezy
		
				break;
			case 2:
				warunek = false;
				System.out.println("Podaj wielkosc danych: ");
				N = input.nextInt();
				// TODO zrob co nalezy
				test(N,false);
				break;
			default:
				System.out.println("Zly wybor spr�boj jeszcze raz");
			}
		}
		
		
		
//		insertionSort(t,t.length);
//		try {
//			t = mergeSort(t,t.length);
//		} catch (Exception e) {
//			// TODO Auto-generated catch block
//			e.printStackTrace();
//		}
//		quickSort(t,0,t.length);
		//int e = partition(t,0,t.length);
		//System.out.println(e);
//		drukuj(t);
		input.close();
	}
	/**
	 * 
	 * @param N size of entry data
	 */
	public static void test(int N, boolean czyRand) {
		int compareKeys = 0;
		int swapKeys = 0;
		for(int i = 0; i<N; i += 100) {	
			int j = 0;
			//wykonujemy 1000 razy dla danego i
			for(j = 0; j < 1000; j++) {
				if(czyRand) {
					t = random(i);
				} else {
					try {
						t = desc(i);
					} catch (Exception e) {
						// TODO Auto-generated catch block
						e.printStackTrace();
					}
				}
				//TODO sort array t and get average info
				insertionSort(t,i);
				compareKeys += insertionCompareKeys;
				swapKeys += insertionSwapKeys;
			}
			compareKeys /= 1000;
			swapKeys /= 1000;
			System.out.println("Insertion sort Dla danych rozmiaru: "+i+" compare: "+compareKeys+" swap: "+swapKeys );
			for(j = 0; j < 1000; j++) {
				if(czyRand) {
					t = random(i);
				} else {		
					t = desc(i);	
				}
				//TODO sort array t and get average info
				mergeCompareKeys = 0;
				mergeSwapKeys = 0;
				mergeSort(t,i);
				compareKeys += mergeCompareKeys;
				swapKeys += mergeSwapKeys;
			}
			compareKeys /= 1000;
			swapKeys /= 1000;
			System.out.println("Merge sort Dla danych rozmiaru: "+i+" compare: "+compareKeys+" swap: "+swapKeys );
			for(j = 0; j < 1000; j++) {
				if(czyRand) {
					t = random(i);
				} else {		
					t = desc(i);	
				}
				//TODO sort array t and get average info
				quickCompareKeys = 0;
				quickSwapKeys = 0;
				quickSort(t,0,i);
				compareKeys += quickCompareKeys;
				swapKeys += quickSwapKeys;
			}
			compareKeys /= 1000;
			swapKeys /= 1000;
			System.out.println("Quick sort Dla danych rozmiaru: "+i+" compare: "+compareKeys+" swap: "+swapKeys );
		}
		
	}
	
	
}
