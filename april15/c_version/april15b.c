// april15b.c


#include <stdio.h>

#define MAX_SSN (1000*1000*1000)

char taxpayerW2[MAX_SSN];
char taxpayerW3[MAX_SSN];
char taxpayerW4[MAX_SSN];
char taxpayerW5[MAX_SSN];
//char taxpayerW6[MAX_SSN];

int main(int argc, char * argv[]) {
	
	printf("april15b.c starting\n");
		for (int i=0; i<MAX_SSN; i++) {
		taxpayerW2[i] = 2;
		taxpayerW3[i] = 3;
		taxpayerW4[i] = 4;
		taxpayerW5[i] = 5;
//		taxpayerW6[i] = 6;
		}

	printf("april15b.c ending\n");
}

