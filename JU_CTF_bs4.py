#!/usr/bin/env python3
import sys
import random
import string
import requests
from multiprocessing import Process

# Change this value only
form_id = '1FAIpQLSegHaWfpaYRlhmcnKs_vQRyZCKoVLLKcw16DAA-l8i7CyJtwg'

url = 'https://docs.google.com/forms/d/e/' + form_id + '/formResponse'

def name():
    return random.choice(string.ascii_uppercase) + text(random.randint(3, 5)) + ' ' + random.choice(string.ascii_uppercase) + text(random.randint(3, 5))

def text(length=5):
    letters = string.ascii_lowercase
    return ''.join(random.choice(letters) for i in range(length))

def phone():
    return '+8801' + str(random.randint(555555555, 999999999))

def uni_name():
    name = ['University of Dhaka',
            'University of Rajshahi',
            'Bangladesh Agricultural University',
            'Bangladesh University of Engineering & Technology',
            'University of Chittagong',
            'Jahangirnagar University',
            'Islamic University, Bangladesh',
            'Shahjalal University of Science and Technology',
            'Khulna University',
            'Bangabandhu Sheikh Mujib Medical University',
            'Bangabandhu Sheikh Mujibur Rahman Agricultural University',
            'Hajee Mohammad Danesh Science & Technology University',
            'Mawlana Bhashani Science and Technology University',
            'Patuakhali Science and Technology University',
            'Sher-e-Bangla Agricultural University',
            'Noakhali Science and Technology University',
            'Chittagong University of Engineering & Technology',
            'Rajshahi University of Engineering & Technology',
            'Khulna University of Engineering & Technology',
            'Dhaka University of Engineering & Technology',
            'Jagannath University',
            'Jatiya Kabi Kazi Nazrul Islam University',
            'Chittagong Veterinary and Animal Sciences University',
            'Sylhet Agricultural University',
            'Comilla University',
            'Jessore University of Science & Technology',
            'Pabna University of Science and Technology',
            'Bangladesh University of Professionals',
            'Begum Rokeya University',
            'Bangladesh University of Textiles',
            'University of Barisal',
            'Bangabandhu Sheikh Mujibur Rahman Science and Technology University',
            'Rangamati Science and Technology University',
            'Bangabandhu Sheikh Mujibur Rahman Maritime University',
            'Rajshahi Medical University',
            'Chittagong Medical University',
            'Rabindra University, Bangladesh',
            'Bangabandhu Sheikh Mujibur Rahman Digital University',
            'Sheikh Hasina University',
            'Sheikh Fazilatunnesa Mujib University of Science & Technology',
            'Sylhet Medical University',
            'Khulna Agricultural University',
            'Bangabandhu Sheikh Mujibur Rahman Aviation and Aerospace University',
            'Islamic University of Technology',
            'Asian University for Women',
            'Bangladesh Open University',
            'National University of Bangladesh',
            'Islamic Arabic University',
            'International University of Business Agriculture and Technology',
            'North South University',
            'University of Science & Technology Chittagong',
            'Central Women\'s University',
            'Independent University, Bangladesh',
            'American International University-Bangladesh',
            'Ahsanullah University of Science and Technology',
            'Dhaka International University',
            'The University of Comilla, Bangladesh',
            'Asian University of Bangladesh',
            'East West University',
            'Gono Bishwabidyalay',
            'People\'s University of Bangladesh',
            'Queens University',
            'University of Asia Pacific (Bangladesh)',
            'BGMEA University of Fashion & Technology',
            'Chittagong Independent University (CIU)',
            'Bangladesh University',
            'BGC Trust University Bangladesh',
            'BRAC University',
            'Manarat International University',
            'Premier University, Chittagong',
            'Pundra University of Science and Technology',
            'Southern University, Bangladesh',
            'Sylhet International University',
            'City University, Bangladesh',
            'Daffodil International University',
            'Green University of Bangladesh',
            'IBAIS University',
            'I Love You Tisha'
    ]

    return random.choice(name)

def trxid():
    tr = 'TrxID ' + text(random.randint(1, 2)) + \
        str(random.randint(1, 9999)) + \
        text(random.randint(1, 2)) + \
        str(random.randint(1, 9999)) + \
        text(random.randint(1, 2)) + \
        str(random.randint(1, 9999)) + \
        text(random.randint(1, 2)) + str(random.randint(1, 9999))

    return tr[:16].upper()

def data():
    form_data = {
        'entry.838304731': name(),
        'entry.2146648422': 2,
        'entry.2129567656': text() + '@' + random.choice(['gmail', 'yahoo', 'outlook', 'yandex', 'hotmail']) + '.com',
        'entry.1114548755': name(),
        'entry.773910481': 'sem ' + str(random.randint(1, 8)),
        'entry.415304891': uni_name(),
        'entry.1184928948': phone(),
        'entry.56430144': random.choice(['M', 'L', 'XL', 'XXL']),
        'entry.2114743164': text() + '@' + random.choice(['gmail', 'yahoo', 'outlook', 'yandex', 'hotmail']) + '.com',
        'entry.1738188966': name(),
        'entry.373845006': 'sem ' + str(random.randint(1, 8)),
        'entry.1791490930': uni_name(),
        'entry.1896004854': phone(),
        'entry.2001288836': random.choice(['M', 'L', 'XL', 'XXL']),
        'entry.1288976591': phone(),
        'entry.1617483818': trxid(),
        'entry.1206724994': random.randint(1, 3),
        'fvv': 1,
        'draftResponse': ['', '', '-3928320261974900822'],
        'pageHistory': 0
    }

    return form_data

def main(target = 1):
    # print(form_data)
    agent = "Mozilla/5.0 (X11; Linux x86_64; rv:60.0) Gecko/20100101 Firefox/60.0"
    user_agent = {'Referer': 'https://docs.google.com/forms/d/e/' + form_id + '/viewform',
                'User-Agent': agent}

    for i in range(target):
        r = requests.post(url, data = data(), headers = user_agent)
        print(f'[+] {i+1}-> Result -> ', r.status_code)

if __name__ == '__main__':
    # print(uni_name())
    # target = int(sys.argv[1])
    # main(target)
    try:
        target = int(sys.argv[1])
        # main(target)
        ps1 = Process(target=main, args=(target,)) # Initiate Process 1
        ps2 = Process(target=main, args=(target,))  # Initiate Process 2

        # Start Process 1
        ps1.start()
        # Start Process 2
        ps2.start()

        # Wait for process
        ps1.join()
        ps2.join()
    except IndexError:
        print('[*] Please provide your target number of sumissions as argument...')
    except KeyboardInterrupt:
        print('\n[*] Yes sir!\n[*] I\'m Exiting...')
    except Exception as e:
        print(e)
