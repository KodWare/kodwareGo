from ctypes import cdll
import time

lib = cdll.LoadLibrary("./lib")
lib.HandleSignals() # :)

while 1:
    print("Birşeyler yapıyorum")
    time.sleep(10)