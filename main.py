import threading
import time
import ply.lex as lex

# Define tokens
tokens = (
    'LBRACE', 'RBRACE', 'IF', 'WHILE', 'FOR', 'OTHER'
)

# Define token rules
t_LBRACE = r'\{'
t_RBRACE = r'\}'
t_IF = r'if'
t_WHILE = r'while'
t_FOR = r'for'
t_OTHER = r'.'

# Ignore spaces and tabs
t_ignore = ' \t'

def t_newline(t):
    r'\n+'
    t.lexer.lineno += len(t.value)

def t_error(t):
    t.lexer.skip(1)

lexer = lex.lex()

def check_braces(code):
    lexer.input(code)
    stack = []
    for token in lexer:
        if token.type == 'LBRACE':
            stack.append(token)
        elif token.type == 'RBRACE':
            if not stack:
                return False
            stack.pop()
    return not stack

def check_if_statements(code):
    lines = code.split('\n')
    for line in lines:
        trimmed_line = line.strip()
        if trimmed_line.startswith('if'):
            parts = trimmed_line.split()
            if len(parts) < 2 or parts[1] == '{':
                return False
    return True

def check_while_statements(code):
    lines = code.split('\n')
    for line in lines:
        trimmed_line = line.strip()
        if trimmed_line.startswith('while'):
            parts = trimmed_line.split()
            if len(parts) < 2 or parts[1] == '{':
                return False
    return True

def check_for_statements(code):
    lines = code.split('\n')
    for line in lines:
        trimmed_line = line.strip()
        if trimmed_line.startswith('for'):
            parts = trimmed_line[3:].split(';')
            if len(parts) != 3:
                return False
    return True

with open('sample.txt', 'r') as file:
    code = file.read()

brace_match = False
if_statements_correct = False
while_statements_correct = False
for_statements_correct = False

start = time.time()

def check_braces_thread():
    global brace_match
    brace_match = check_braces(code)
    if brace_match:
        print("Braces matched successfully.")
    else:
        print("Braces do not match, exiting code.")

def check_if_statements_thread():
    global if_statements_correct
    if_statements_correct = check_if_statements(code)
    if if_statements_correct:
        print("'If' statements are correct.")
    else:
        print("'If' statements are incorrect, exiting code.")

def check_while_statements_thread():
    global while_statements_correct
    while_statements_correct = check_while_statements(code)
    if while_statements_correct:
        print("'While' statements are correct.")
    else:
        print("'While' statements are incorrect, exiting code.")

def check_for_statements_thread():
    global for_statements_correct
    for_statements_correct = check_for_statements(code)
    if for_statements_correct:
        print("'For' statements are correct.")
    else:
        print("'For' statements are incorrect, exiting code.")

threads = [
    threading.Thread(target=check_braces_thread),
    threading.Thread(target=check_if_statements_thread),
    threading.Thread(target=check_while_statements_thread),
    threading.Thread(target=check_for_statements_thread)
]

for thread in threads:
    thread.start()

for thread in threads:
    thread.join()

if not (brace_match and if_statements_correct and while_statements_correct and for_statements_correct):
    elapsed = time.time() - start
    print(f"Finished checks in {elapsed:.2f} seconds")
    print("--------------------------------")
    print("Code is invalid.")
    exit()

elapsed = time.time() - start
print(f"Finished checks in {elapsed:.2f} seconds")
print("--------------------------------")
print("Code is valid.")
