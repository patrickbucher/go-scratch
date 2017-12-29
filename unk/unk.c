#include <stdio.h>
#include <string.h>

int main()
{
    char c;
    char pattern[] = "<unk>";
    char replacement[] = "<raw_unk>"; 
    char buf[5]; // length of "<unk>"
    int i, j, pat_len;

    pat_len = strlen(pattern);
    for (i = 0; (c = getchar()) != EOF;) {
        if (c == pattern[i]) {
            buf[i++] = c;
            if (i == pat_len) {
                printf("%s", replacement);
                i = 0;
            }
        } else {
            if (i > 0) {
                for (j = 0; j < i; j++) {
                    putchar(buf[j]);
                }
                i = 0;
            }
            putchar(c);
        }
    }

    return 0;
}
