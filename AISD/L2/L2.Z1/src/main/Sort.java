package main;

import java.util.Random;
import java.util.Scanner;

public class Sort {
	static int t[];
	/*
	 * Tablica danych t , n to romiar tablicy
	 */
	public static void insertionSort(int t[], int n) {
		if(n == 0) {
			System.out.println("Brak argument�w");
		}
		for (int i = 1; i < n; i++) {
			int key = t[i];
			int j = i-1;
			while(j >= 0 && key < t[j]) {
				t[j+1]=t[j];
				j--;
			}
			t[j+1] = key;
		}
	}
	
	public static int[] mergeSort(int t[], int n) throws Exception {
		if(n == 0) {
			throw new Exception("Brak Argument�w");
			
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
				ret[j+k] = A[j];
				j++;
			} else if (A[j] > B[k]) {
				ret[j+k] = B[k];
				k++;
			} else {
				ret[j+k] = A[j];
				j++;
			}
		}
		
		// wiemy ze pierwsza tablica juz skonczyla a w drugiej sa same wieksze elementy
		
		while(k < dlB) {
			ret[j+k] = B[k];
			k++;
		}
	
		while(j < dlA) {
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
		int i = p, tmp;
		for(int j=p+1; j<q; j++) {
			if(A[j] <= pivot) {
				
				//swap A[j],a[i+1]
				i++;
				tmp = A[j];
				A[j] = A[i];
				A[i] = tmp;		
			}
		}
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
	
	public static void random(int t[]) {
		
		Random r = new Random();
		
		for(int i = 0; i< t.length; i++) {
			t[i] = r.nextInt();
		}
		//drukuj(t);		
	}
	
	public static void desc(int t[]) throws Exception{
		Random r = new Random();
		if(t.length == 0) {
			throw new Exception("Brak argument�w");
		} else {
			t[0] = r.nextInt();
			for(int i = 1; i < t.length; i++){
				t[i] = t[i-1]-2;
			}
		}
		//drukuj(t);
	}
	
	public static void main(String[] args) {
		
		
		int decision = 0;
		int N = 0;
		boolean warunek = true;
		
		Scanner input = new Scanner(System.in);
		System.out.println("Wybierz odpowiednia liczbe: ");
		System.out.println("1. Losowy ciag");
		System.out.println("2. Ciag malejacy");
		decision = input.nextInt();
		while(warunek) {
			switch(decision) {
			case 1:
				warunek = false;
				System.out.println("Podaj wielkosc danych: ");
				N = input.nextInt();
				t = new int[N];
				// TODO zrob co nalezy
				random(t);
				break;
			case 2:
				warunek = false;
				System.out.println("Podaj wielkosc danych: ");
				N = input.nextInt();
				// TODO zrob co nalezy
				t = new int[N];
				try {
					desc(t);
				} catch (Exception e) {
					// TODO Auto-generated catch block
					e.printStackTrace();
				}
				break;
			default:
				System.out.println("Zly wybor spr�boj jeszcze raz");
			}
		}
		
//		insertionSort(t,t.length);
		try {
			t = mergeSort(t,t.length);
		} catch (Exception e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
//		quickSort(t,0,t.length);
		//int e = partition(t,0,t.length);
		//System.out.println(e);
		drukuj(t);
	}
}