import ctypes
lib = ctypes.cdll.LoadLibrary("./sözDizimi") # GO ile oluşturduğumuz kütüphaneyi Py'a import ettik
lib.Merhaba()
#lib.PrintUn("100") #  Stringler bozuluyor :)

print("return olan veri alınıyor : ")
print(lib.MerhabaUn()) # 687522154 // Biz string döndürdük ama bu rakam döndürdü ¿? :)