import PIL.Image

im = PIL.Image.open('grape.png')
im_resized = PIL.Image.new(im.mode, (162*2, 179*2))
for r in range(im_resized.size[1]):
    for c in range(im_resized.size[0]):
        rr = round((r + 1) / im_resized.size[1] * im.size[1]) - 1
        cc = round((c + 1) / im_resized.size[0] * im.size[0]) - 1
        im_resized.putpixel((c, r), im.getpixel((cc, rr)))
im_resized.show()
