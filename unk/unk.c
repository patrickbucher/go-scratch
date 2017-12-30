#include <stdio.h>
#include <string.h>

#define PAT_LEN 5

int main()
{
    char pattern[] = "<unk>";
    char replacement[] = "<raw_unk>"; 
    int c;
    int i, j;

    for (i = 0; (c = getchar()) != EOF;) {
        if (c == pattern[i]) {
            i++;
            if (i == PAT_LEN) {
                printf("%s", replacement);
                i = 0;
            }
        } else {
            if (i > 0) {
                for (j = 0; j < i; j++) {
                    putchar(pattern[j]);
                }
                i = 0;
            }
            if (c == pattern[0]) {
                i = 1;
            } else {
                putchar(c);
            }
        }
    }

    return 0;
}
