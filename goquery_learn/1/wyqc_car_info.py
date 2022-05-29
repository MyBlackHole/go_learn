from update_jieba_key import jieba_sort_md5
from lxml import etree
import requests
import json
import hashlib
import pymysql
import jieba
import re
# from zdy import get_zdy_proxy
from logger import logger
from pathlib import Path
from multiprocessing import Pool


info_dict = {
    'brand':'',
    'model':'',
    'style':'',
    'level':'',
    'engine':'',
    'car_length':'',
    'car_width':'',
    'car_height':'',
    'wheelbase':'',
    'front_braking':'',
    'rear_braking':'',
    'tires_count':4,
    'front_tires':'',
    'rear_tires':'',
}

def hash_md5(text: str) -> str:
    """md5"""
    m = hashlib.md5()
    s_bytes = text.encode().decode().encode()
    m.update(s_bytes)
    return m.hexdigest()

def insert_db(value):
    try:
        conn = pymysql.connect(host='192.168.1.111', port=3383, user='yunrun', passwd='yr2020!QAZ', db='urun_car')
        cursor = conn.cursor()
        sql = "INSERT INTO car (`site`,`brand`,`model`,`style`,`level`,`engine`,`car_length`,`car_width`,`car_height`,`wheelbase`,`front_braking`,`rear_braking`,`tires_count`,`front_tires`,`rear_tires`,`source`,`jieba_sort_md5`) values (%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)"
        cursor.execute(sql, value)
        conn.commit()
        cursor.close()
        conn.close()
        # logger.info(f'{value}入库成功')
        logger.info(f'{value}入库成功')
        return cursor.lastrowid
    except Exception as e:
        logger.error(f'在car的表插入数据{value}时出错了,{e}')


def get_info(path_name):
    with open(path_name,'r',encoding='utf-8') as f:
        html2 = f.read()
    tree = etree.HTML(html2)
    data, err = get_dict(tree,path_name)
    if err:
        logger.debug(f"{data}, {err}")
        return data, err

    if data:
        if (data.get('brand'),data['model'],data['style']) not in car_db_info:
            jbmd5_name = data.get('brand') + data['model'] + data['style']
            jbmd5 = jieba_sort_md5(jbmd5_name)
            value = (1, data.get('brand'), data['model'], data['style'], data['level'], data['engine'], data['car_length'],data['car_width'], data['car_height'], data['wheelbase'], data['front_braking'], data['rear_braking'],data['tires_count'], data['front_tires'], data['rear_tires'], 111, jbmd5)
            data_id = insert_db(value)
            data['data_id'] = data_id
            name = './car_info/' + str(path_name.stem) + '.json'
            with open(name,'w',encoding='utf-8') as f:
                f.write(json.dumps(data,ensure_ascii=False))
        else:
            return data, f'{path_name}已经存在'


def get_dict(tree,path_name):
    data = info_dict.copy()
    name_info = ""
    try:
        name_info = tree.xpath('//meta[@name="keywords"]/@content')
    except Exception as e1:
        return str(path_name), str(e1)
    if name_info:
        name_info = name_info[0]
        brand = name_info.split(',')[0]
        model = name_info.split(',')[1]
        style = name_info.split(',')[2]
        try:
            # 发动机
            engine = tree.xpath('//*[@id="container-main"]/div[1]/div[2]/div[2]/div[1]/div[2]/div[2]/table/tbody/tr[2]/td[2]/span//text()')[0]
            # 级别
            level = ""
            car_szie = tree.xpath('//*[@id="container-main"]/div[1]/div[2]/div[2]/div[1]/div[2]/div[2]/table/tbody/tr[1]/td[6]/span//text()')[0]
            car_length = car_szie.split('*')[0]
            car_width = car_szie.split('*')[1]
            car_height = car_szie.split('*')[2]

            wheelbase = tree.xpath('//td/span[@data-title="轴距(mm)"]/../following-sibling::td//text()')[1]
            front_braking = tree.xpath('//td/span[@data-title="前制动类型"]/../following-sibling::td//text()')[1]
            rear_braking = tree.xpath('//td/span[@data-title="后制动类型"]/../following-sibling::td//text()')[1]
            front_tires = tree.xpath('//td/span[@data-title="前轮胎规格"]/../following-sibling::td//text()')[1]
            rear_tires = tree.xpath('//td/span[@data-title="后轮胎规格"]/../following-sibling::td//text()')[1]
            data['brand'] = brand
            data['model'] = model
            data['style'] = style
            data['level'] = level
            data['engine'] = engine
            data['car_length'] = car_length
            data['car_width'] = car_width
            data['car_height'] = car_height
            data['wheelbase'] = wheelbase
            data['front_braking'] = front_braking
            data['rear_braking'] = rear_braking
            data['front_tires'] = front_tires
            data['rear_tires'] = rear_tires
            data['path_name'] = str(path_name)
            return data, None
        except Exception as e:
            return str(path_name), str(e)
    else:
        return str(path), str(f'{path_name}文件拿不到车辆名字')


def select_db_car():
    conn = pymysql.connect(host='192.168.1.111', port=3383, user='yunrun', passwd='yr2020!QAZ', db='urun_car')
    cursor = conn.cursor()
    sql = "SELECT `brand`,`model`,`style`  FROM car WHERE SOURCE=111"
    cursor.execute(sql)
    info = cursor.fetchall()
    with open('./yicun_car.json', 'w',encoding='utf-8') as f:
        f.write(json.dumps(info,ensure_ascii=False))
    return info

def read_info():
    with open('./yicun_car.json', 'r',encoding='utf-8') as f:
        car_db_info = json.loads(f.read())
    return car_db_info


if __name__ == '__main__':
    # 解析入库
    car_db_info = select_db_car()

    path =Path('product.auto.163.com')
    product = path / 'product'
    # path = 'C:\各大汽车网站\网易\\product'
    # paths = get_doc_name(path)

    p = Path('bu_car_img_send_log')
    debug_log = p / 'debug_log.log'

    key_list = []
    resp_list = re.findall(r"car_info/(.*)\.json", debug_log.read_text())

    # with open('key_list', 'r') as f:
        # text = f.read()
    path_list = []
    for item in resp_list:
        name = item + ".html"
        path_list.append(product / name) 

    with Pool(1) as p:
        resp_list = p.map(get_info, path_list)


    # with open('./cs.txt', 'r', encoding='utf-8') as f:
    #     info = f.read()
    # data = re.findall('- (.*?)文件解析的时候报错了', info)
    # for path_name in data:
    #     print(path_name)
    #     get_info(path_name)


    # path = 'C:\各大汽车网站\网易\product.auto.163.com\product\\000CLfAP.html'
    # get_info(path)
    # path_name = './000BAYDZ.html'
    # get_info(path_name)

    # # 充数据
    # p = Path('car_info_log')
    # err_log = p / 'err_log.log' 
    # text = err_log.read_text()
    # p_list = re.findall(r'\((.*)\)', text)
    # for item in p_list:
    #     value = item.split(',')
    #     data_id = insert_db(value)
    #     data['data_id'] = data_id
    #     name = './car_info/' + str(path_name.stem) + '.json'
    #     with open(name,'w',encoding='utf-8') as f:
    #         f.write(json.dumps(data,ensure_ascii=False))

