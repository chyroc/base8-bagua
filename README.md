# base8-bagua

八卦图形base8

## install
```bash
go get -u github.com/Chyroc/base8-bagua/cmd/...
```

## usage

```bash
base8-bagua -h
NAME
     base8-bagua -- Encode and decode using Base8-Bagua representation

SYNOPSIS
     base8 [-D|-h]
```

```bash
echo 'hello base8-bagua' | base8-bagua
☳☲☰☶☲☵☵☴☳☳☰☶☷☴☴☰☳☰☴☶☰☵☶☳☳☱☲☳☴☰☵☵☳☰☴☶☰☵☴☷☳☵☲☶☰☴

echo '☳☲☰☶☲☵☵☴☳☳☰☶☷☴☴☰☳☰☴☶☰☵☶☳☳☱☲☳☴☰☵☵☳☰☴☶☰☵☴☷☳☵☲☶☰☴' | base8-bagua -D
hello base8-bagua
```

## 解释

### 符号对应数字

| 符号 | 卦名 | 拼音 | 8进制数 |
| :--: | :--: | :--: | :-----: |
|  ☰   |  乾  | qián |    0    |
|  ☱   |  兑  | duì  |    1    |
|  ☲   |  离  |  lí  |    2    |
|  ☳   |  震  | zhèn |    3    |
|  ☴   |  巽  | xùn  |    4    |
|  ☵   |  坎  | kǎn  |    5    |
|  ☶   |  艮  | gèn  |    6    |
|  ☷   |  坤  | kūn  |    7    |

### encode

* string转成byte数组
* byte转二进制，拼接byte数组为二进制数组
* 3个二进制数为一个单位，划分成n份（最后补0）
* 将每个单位转化成8进制数字（3位二进制范围：0-7）
* 根据上表的对应关系，换成八卦符号
* 返回八卦符号字符串

### decode

* 将八卦符号字符串转成8进制数数组
* 将8进制数字转换成3个二进制数，组成二进制数组
* 以8个二进制数为一个byte，组成byte数组
* 返回byte数组
