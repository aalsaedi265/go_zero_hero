from flask import Flask, render_template, request

app = Flask(__name__, template_folder='templates')

ITEMS = list(range(1,101))

@app.route('/')

def index():
    page = request.args.get('page', 1, type=int)

    per_page =10
    start = (page -1) *per_page
    end = start + per_page
    total_page = (len(ITEMS) + per_page-1 )// per_page
    
    items_page = ITEMS[start:end]
    
    return render_template(template_name_or_list='index.html', items_page= items_page, total_page = total_page, page = page)

if __name__ == '__main__':
    app.run(debug=True)