#!/usr/bin/env python3
"""Generate Windows AMD64 PE icon resources from OpenCompanionListener.ico.

Only needed when replacing the application icon. Standard Windows builds use
listener/icon_windows_amd64.syso as already checked in, with no Python needed.
"""
from pathlib import Path
import struct
p=Path(__file__).resolve().parents[1]
b=(p/'OpenCompanionListener.ico').read_bytes()
reserved,itype,count=struct.unpack_from('<HHH',b,0)
assert (reserved,itype)==(0,1)
icons=[]
for i in range(count):
    w,h,cc,rs,planes,bits,length,offset=struct.unpack_from('<BBBBHHII',b,6+16*i)
    icons.append((w,h,cc,planes,bits,b[offset:offset+length]))
# Named resources: RT_ICON/each ID and RT_GROUP_ICON/ID1
res={3:{i+1:{0x409:image[-1]} for i,image in enumerate(icons)},14:{1:{0x409:b''}}}
group=struct.pack('<HHH',0,1,count)
for n,(w,h,cc,planes,bits,data) in enumerate(icons,1):
    group+=struct.pack('<BBBBHHIH',w,h,cc,0,planes,bits,len(data),n)
res[14][1][0x409]=group
# alloc directories in raw section
raw=bytearray()
def alloc(sz):
    while len(raw)%4:raw.append(0)
    at=len(raw);raw.extend(b'\0'*sz);return at
def put(off,fmt,*vals):struct.pack_into(fmt,raw,off,*vals)
root=alloc(16+8*len(res))
type_dirs={}; name_dirs={}; entries={}; data_blobs={}
for tid,ids in res.items():
    type_dirs[tid]=alloc(16+8*len(ids))
    for rid,langs in ids.items():
        name_dirs[tid,rid]=alloc(16+8*len(langs))
        for lang,data in langs.items():entries[tid,rid,lang]=alloc(16)
for tid,ids in res.items():
    for rid,langs in ids.items():
        for lang,data in langs.items():
            at=alloc(len(data));raw[at:at+len(data)]=data
            data_blobs[tid,rid,lang]=(at,len(data))
# dir headers and entries
put(root,'<IIHHHH',0,0,0,0,0,len(res))
for idx,(tid,ids) in enumerate(sorted(res.items())):
    put(root+16+idx*8,'<II',tid,type_dirs[tid]|0x80000000)
    td=type_dirs[tid];put(td,'<IIHHHH',0,0,0,0,0,len(ids))
    for j,(rid,langs) in enumerate(sorted(ids.items())):
        put(td+16+j*8,'<II',rid,name_dirs[tid,rid]|0x80000000)
        nd=name_dirs[tid,rid];put(nd,'<IIHHHH',0,0,0,0,0,len(langs))
        for k,lang in enumerate(sorted(langs)):
            put(nd+16+k*8,'<II',lang,entries[tid,rid,lang])
            de=entries[tid,rid,lang];at,size=data_blobs[tid,rid,lang]
            put(de,'<IIII',at,size,0,0)
# IMAGE_REL_AMD64_ADDR32NB relative to rsrc symbol
relocations=b''.join(struct.pack('<IIH',entry,0,3) for entry in entries.values())
# COFF header + section header + raw + relocations + symbol table + string table
ptr_raw=60
ptr_rel=ptr_raw+len(raw)
ptr_sym=ptr_rel+len(relocations)
head=struct.pack('<HHIIIHH',0x8664,1,0,ptr_sym,1,0,0)
section=struct.pack('<8sIIIIIIHHI',b'.rsrc\0\0\0',0,0,len(raw),ptr_raw,ptr_rel,0,len(entries),0,0x40000040)
sym=struct.pack('<8sIhHBB',b'.rsrc\0\0\0',0,1,0,3,0)
payload=head+section+bytes(raw)+relocations+sym+struct.pack('<I',4)
out=p/'listener/icon_windows_amd64.syso';out.write_bytes(payload)
print('icon syso bytes',len(payload),'resource bytes',len(raw),'relocs',len(entries))
