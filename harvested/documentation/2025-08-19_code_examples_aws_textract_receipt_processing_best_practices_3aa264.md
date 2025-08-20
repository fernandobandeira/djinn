---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T20:31:12.908961'
profile: code_examples
source: https://aws.amazon.com/blogs/machine-learning/automatically-extract-text-and-structured-data-from-documents-with-amazon-textract/
topic: AWS Textract Receipt Processing Best Practices
---

# AWS Textract Receipt Processing Best Practices

## Select your cookie preferences
We use essential cookies and similar tools that are necessary to provide our site and services. We use performance cookies to collect anonymous statistics, so we can understand how customers use our site and make improvements. Essential cookies cannot be deactivated, but you can choose “Customize” or “Decline” to decline performance cookies. If you agree, AWS and approved third parties will also use cookies to provide useful site features, remember your preferences, and display relevant content, including relevant advertising. To accept or decline all non-essential cookies, choose “Accept” or “Decline.” To make more detailed choices, choose “Customize.”
AcceptDeclineCustomize
## Customize cookie preferences
We use cookies and similar tools (collectively, "cookies") for the following purposes.
### Essential
Essential cookies are necessary to provide our site and services and cannot be deactivated. They are usually set in response to your actions on the site, such as setting your privacy preferences, signing in, or filling in forms. 
### Performance
Performance cookies provide anonymous statistics about how customers navigate our site so we can improve site experience and performance. Approved third parties may perform analytics on our behalf, but they cannot use the data for their own purposes.
Allowed
### Functional
Functional cookies help us provide useful site features, remember your preferences, and display relevant content. Approved third parties may set these cookies to provide certain site features. If you do not allow these cookies, then some or all of these services may not function properly.
Allowed
### Advertising
Advertising cookies may be set through our site by us or our advertising partners and help us deliver relevant marketing content. If you do not allow these cookies, you will experience less relevant advertising.
Allowed
Blocking some types of cookies may impact your experience of our sites. You may review and change your choices at any time by selecting Cookie preferences in the footer of this site. We and selected third-parties use cookies or similar technologies as specified in the [AWS Cookie Notice](https://aws.amazon.com/legal/cookies/).
CancelSave preferences
## Your privacy choices
We and our advertising partners (“we”) may use information we collect from or about you to show you ads on other websites and online services. Under certain laws, this activity is referred to as “cross-context behavioral advertising” or “targeted advertising.”
To opt out of our use of cookies or similar technologies to engage in these activities, select “Opt out of cross-context behavioral ads” and “Save preferences” below. If you clear your browser cookies or visit this site from a different device or browser, you will need to make your selection again. For more information about cookies and how we use them, read our [Cookie Notice](https://aws.amazon.com/legal/cookies/).
Allow cross-context behavioral adsOpt out of cross-context behavioral ads
To opt out of the use of other identifiers, such as contact information, for these activities, fill out the form [here](https://pulse.aws/application/ZRPLWLL6?p=0).
For more information about how AWS handles your information, read the [AWS Privacy Notice](https://aws.amazon.com/privacy/).
CancelSave preferences
## Unable to save cookie preferences
We will only store essential cookies at this time, because we were unable to save your cookie preferences.If you want to change your cookie preferences, try again later using the link in the AWS console footer, or contact support if the problem persists.
Dismiss
[ Skip to Main Content](https://aws.amazon.com/blogs/machine-learning/automatically-extract-text-and-structured-data-from-documents-with-amazon-textract/#aws-page-content-main)
  * English
  * [Contact us](https://aws.amazon.com/contact-us/?nc2=h_ut_cu)
  * Support 
  * My account 


  * [](https://aws.amazon.com/?nc2=h_home)
  * [Agentic AI](https://aws.amazon.com/ai/agentic-ai/?nc2=h_l1_f)
  * Discover AWS
  * Products
  * Solutions
  * Pricing
  * More 


  * Filter: All
  * Sign in to console
  * [Create Account](https://portal.aws.amazon.com/gp/aws/developer/registration/?nc2=h_su&src=header_signup)


AWS Blogs
  * [Home](https://aws.amazon.com/blogs/)
  * Blogs
  * Editions


## [Artificial Intelligence](https://aws.amazon.com/blogs/machine-learning/)
# Automatically extract text and structured data from documents with Amazon Textract
by Kashif Imran, Sameer Shrivastava, Martin Schade, and Urvi Sharma on 30 MAY 2019 in [Amazon Comprehend](https://aws.amazon.com/blogs/machine-learning/category/artificial-intelligence/amazon-comprehend/ "View all posts in Amazon Comprehend"), [Amazon Machine Learning](https://aws.amazon.com/blogs/machine-learning/category/artificial-intelligence/amazon-machine-learning/ "View all posts in Amazon Machine Learning"), [Amazon Textract](https://aws.amazon.com/blogs/machine-learning/category/artificial-intelligence/amazon-textract/ "View all posts in Amazon Textract"), [Amazon Translate](https://aws.amazon.com/blogs/machine-learning/category/artificial-intelligence/amazon-translate/ "View all posts in Amazon Translate"), [Artificial Intelligence](https://aws.amazon.com/blogs/machine-learning/category/artificial-intelligence/ "View all posts in Artificial Intelligence"), [Top Posts](https://aws.amazon.com/blogs/machine-learning/category/top-posts/ "View all posts in Top Posts") [Permalink](https://aws.amazon.com/blogs/machine-learning/automatically-extract-text-and-structured-data-from-documents-with-amazon-textract/) [ Comments](https://aws.amazon.com/blogs/machine-learning/automatically-extract-text-and-structured-data-from-documents-with-amazon-textract/#Comments) [ Share](https://aws.amazon.com/blogs/machine-learning/automatically-extract-text-and-structured-data-from-documents-with-amazon-textract/)
  * [](https://www.facebook.com/sharer/sharer.php?u=https://aws.amazon.com/blogs/machine-learning/automatically-extract-text-and-structured-data-from-documents-with-amazon-textract/)
  * [](https://twitter.com/intent/tweet/?text=Automatically%20extract%20text%20and%20structured%20data%20from%20documents%20with%20Amazon%20Textract&via=awscloud&url=https://aws.amazon.com/blogs/machine-learning/automatically-extract-text-and-structured-data-from-documents-with-amazon-textract/)
  * [](https://www.linkedin.com/shareArticle?mini=true&title=Automatically%20extract%20text%20and%20structured%20data%20from%20documents%20with%20Amazon%20Textract&source=Amazon%20Web%20Services&url=https://aws.amazon.com/blogs/machine-learning/automatically-extract-text-and-structured-data-from-documents-with-amazon-textract/)
  * 

**_July 2025: Post was reviewed and updated for accuracy._**
Documents are a primary tool for record keeping, communication, collaboration, and transactions across many industries, including financial, medical, legal, and real estate. The millions of mortgage applications and hundreds of millions of W2 tax forms processed each year are just a few examples of such documents. A lot of information is locked in unstructured documents. It usually requires time-consuming and complex processes to enable search and discovery, business process automation, and compliance control for these documents.
In this post, we show how you can take advantage of [Amazon Textract](https://aws.amazon.com/textract) to automatically extract text and data from scanned documents without machine learning (ML) experience. While AWS takes care of building, training, and deploying advanced ML models in a highly available and scalable environment, you take advantage of these models with simple-to-use API actions. We cover the following use cases in this post:
  * Text detection from documents
  * Form and table extraction and processing
  * Extract information from identity documents
  * Extract information from invoices and receipts
  * Multi-column detection and reading order
  * Natural language processing and document classification
  * Natural language processing for medical documents
  * Document translation
  * Search and discovery
  * Compliance control with document redaction
  * PDF and multi-page TIFF document processing


## Amazon Textract overview
Before we get started with the use cases, let’s review and introduce some of the core features. Amazon Textract goes beyond simple optical character recognition (OCR) to also identify the contents of fields in forms, information stored in tables, handwritten text, and check boxes. This allows you to use Amazon Textract to instantly read almost any type of document and accurately extract text and data without the need effort or custom code.
The following images show an example document using Amazon Textract on the [AWS Management Console](https://console.aws.amazon.com/textract/) on the **Forms** output tab.
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/06/09/1-4078-Analyze.jpg)
To quickly download a .zip file containing the output, choose **Download results**. You can choose various formats, including raw JSON, text, and CSV files for forms and tables.
In addition to the detected content, Amazon Textract provides additional information like confidence scores and bounded boxes for detected elements. It gives you control of how you consume extracted content and integrate it into various business applications.
Amazon Textract provides both [synchronous](https://docs.aws.amazon.com/textract/latest/dg/sync.html) and [asynchronous](https://docs.aws.amazon.com/textract/latest/dg/async.html) API actions to [extract document text](https://docs.aws.amazon.com/textract/latest/dg/detecting-document-text.html) and [analyze the document text data](https://docs.aws.amazon.com/textract/latest/dg/analyzing-document-text.html). Synchronous APIs can be used for [single-page documents](https://docs.aws.amazon.com/textract/latest/dg/sync.html) and low-latency use cases such as mobile capture. Asynchronous APIs can be used for [multipage documents](https://docs.aws.amazon.com/textract/latest/dg/async.html) such as PDF or TIFF documents with thousands of pages. For more information, see the [Amazon Textract API Reference](https://docs.aws.amazon.com/textract/latest/dg/API_Reference.html).
## Use cases overview
You can take advantage of Amazon Textract API operations using the [AWS SDK](https://aws.amazon.com/tools/) to build power-smart applications. We also use [Amazon Textract Helper](https://github.com/aws-samples/amazon-textract-textractor/tree/master/helper), [Amazon Textract Caller](https://github.com/aws-samples/amazon-textract-textractor/tree/master/caller), [Amazon Textract PrettyPrinter](https://github.com/aws-samples/amazon-textract-textractor/tree/master/prettyprinter), and [Amazon Textract Response Parser](https://github.com/aws-samples/amazon-textract-response-parser) for some of the following use cases. These packages are published to [PyPI](https://pypi.org/user/rekognition-textract-demos/) to speed up development and integration even further.
## Text detection from documents
We start with a simple example of how to detect text from a document. We use the following image as an input document to Amazon Textract. The sample image isn’t good quality, but Amazon Textract can still detect the text with accuracy.
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/06/09/2-4078-address.jpg)
The way to extract information from this document programmatically is through installing [Amazon Textract Helper](https://github.com/aws-samples/amazon-textract-textractor/tree/master/helper):
```
python -m pip install amazon-textract-helper
```

Python
Then we call Amazon Textract to extract information from the document and display the results by running the command line tool:
```
amazon-textract --input-document "s3://amazon-textract-public-content/blogs/amazon-textract-sample-text-amazon-dot-com.png" --pretty-print LINES-
```

Python
The following screenshot shows our output.
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/06/09/3-4078-output.jpg)
The command line tool uses the [Amazon Textract Caller](https://github.com/aws-samples/amazon-textract-textractor/tree/master/caller), [Amazon Textract PrettyPrint](https://github.com/aws-samples/amazon-textract-textractor/tree/master/prettyprinter), and [Amazon Textract Overlayer](https://github.com/aws-samples/amazon-textract-textractor/tree/master/overlayer) packages to generate the results.
The original Amazon Textract response is in JSON format and has the following format:
```
{
  "Blocks": [
    {
      "Geometry": {
        "BoundingBox": {
          "Width": 1.0, 
          "Top": 0.0, 
          "Left": 0.0, 
          "Height": 1.0
        }, 
        "Polygon": [
          {
            "Y": 0.0, 
            "X": 0.0
          }, 
          {
            "Y": 0.0, 
            "X": 1.0
          }, 
          {
            "Y": 1.0, 
            "X": 1.0
          }, 
          {
            "Y": 1.0, 
            "X": 0.0
          }
        ]
      }, 
      "Relationships": [
        {
          "Type": "CHILD", 
          "Ids": [
            "2602b0a6-20e3-4e6e-9e46-3be57fd0844b", 
            "82aedd57-187f-43dd-9eb1-4f312ca30042", 
            "52be1777-53f7-42f6-a7cf-6d09bdc15a30", 
            "7ca7caa6-00ef-4cda-b1aa-5571dfed1a7c"
          ]
        }
      ], 
      "BlockType": "PAGE", 
      "Id": "8136b2dc-37c1-4300-a9da-6ed8b276ea97"
    }..... 
    
  ], 
  "DocumentMetadata": {
    "Pages": 1
  }
}

```

JSON
By using [Amazon Textract Response Parser](https://github.com/aws-samples/amazon-textract-response-parser/blob/master/src-python/README.md), it’s to de-serialize the JSON response and use in your program, the same way [Amazon Textract Helper](https://github.com/aws-samples/amazon-textract-textractor/tree/master/helper) and [Amazon Textract PrettyPrinter](https://github.com/aws-samples/amazon-textract-textractor/tree/master/prettyprinter) use it. The [GitHub repository](https://github.com/aws-samples/amazon-textract-response-parser/blob/master/src-python/README.md#parse-json-response-from-textract) shows some examples.
## Form and table extraction and processing
Amazon Textract can provide the inputs required to automatically process forms and tables without human intervention. For example, a bank could write code to read PDFs of loan applications. The information contained in the document could be used to necessary background and credit checks to approve the loan so that customers can get instant results for their application rather than having to wait several days for manual review and validation.
The following image is an employment application with form fields, check boxes, and a table.
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/06/09/4-4078-application.jpg)
The following code example extracts forms from the employment application and processes different fields:
```
export AWS_DEFAULT_REGION=us-east-2; amazon-textract --input-document "s3://amazon-textract-public-content/blogs/employeeapp20210510.png" --pretty-print FORMS TABLES --features FORMS TABLES
```

Code
The preceding commands produce the following output to visualize the structure of the information.
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/06/09/5-4078-output.jpg)
The key-value pairs from the `FORMS` output are rendered as a table with `Key` and `Value` headlines to allow for processing.
For example, changing the output format by including `—pretty-print-table-format=csv parameter` outputs the data in CSV format (check `amazon-textract —help` for a list of other formats):
```
export AWS_DEFAULT_REGION=us-east-2; amazon-textract --input-document "s3://amazon-textract-public-content/blogs/employeeapp20210510.png" --pretty-print FORMS TABLES --features FORMS TABLES --pretty-print-table-format=csv
```

Python
The following screenshot shows the output.
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/06/09/6-4078-output.jpg)
Amazon Textract can detect tables and their content. A company can amounts from an expense report (as in the following screenshot) and apply rules, such as expense more than $1,000 needs extra review.
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/06/09/7-4078-expense-report.jpg)
The following code uses the CSV output from the command line tool and the sample expense report to print the content of each cell, along with a warning message if expense is more than $1,000:
```
import csv
import sys
from tabulate import tabulate
reader = csv.reader(sys.stdin)
def isFloat(input):
 try:
  float(input)
  return True
 except ValueError:
  return False
all_rows = list()
for row in reader:
  warning = ""
  if len(row)>4:
   if row[4] and isFloat(row[4]):
    if float(row[4]) > 1000.00 and row[3] and not row[3].strip() == 'Total':
     warning = "Warning - value > $1000.00 and requires review."
   row.append(warning)
  all_rows.append(row)
print(tabulate(all_rows, tablefmt='github'))

```

Python
Save this code as `test-csv.py` or copy it from [Amazon Simple Storage Service](http://aws.amazon.com/s3) (Amazon S3) at `s3://amazon-textract-public-content/blogs/test-csv.py`. Then use the following command:
```
export AWS_DEFAULT_REGION=us-east-2; amazon-textract --input-document "s3://amazon-textract-public-content/blogs/expense-report-example.png" --features TABLES --pretty-print TABLES --pretty-print-table-format csv | python test-csv.py
```

Python
We receive the following output.
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/06/09/8-4078-output.jpg)
To recap, we started with a document image, called Amazon Textract to identify and receive the table structure and information, applied business logic on the data, and triggered a business process based on the information.
### Extract information from invoices and receipts
Invoices and receipts are difficult to process at scale because they follow no set design rules, yet any individual customer encounters thousands of distinct types of these documents. The Amazon Textract [AnalyzeExpense](https://docs.aws.amazon.com/textract/latest/dg/API_AnalyzeExpense.html) action identifies standard fields and line-item details for these document types.
The standard fields supported include “Vendor Name”, “Total”, “Receiver Address”, “Invoice/Receipt Date”, “Invoice/Receipt ID”, “Payment Terms”, “Subtotal”, “Due Date”, “Tax”, “Invoice Tax Payer ID”, “Item Name”, “Item Price”, “Item Quantity” plus line-item details. For a complete list check the [Analyzing Invoices and Receipts documentation](https://docs.aws.amazon.com/textract/latest/dg/invoices-receipts.html).
The [AWS Management Console](https://console.aws.amazon.com/textract/) offers options to test the AnalyzeExpense action through the “**Select Document** ” options “**Receipt** ” (image below) or “**Invoice** ” or by “**Choose File** ” option. The latter allows uploading of a document and subsequent selection of “**Analyze Expense** ” in the output tab on the right side. Through “**Download results** ” a zip file including the line-item fields and summary fields can be received.
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/09/10/1-Textract-ReceiptAPI.jpg)
You can call the `AnalyzeExpense` API using the [AWS Command Line Interface](http://aws.amazon.com/cli) (AWS CLI), as shown in the following code. Make sure you have AWS CLI version >= 2.2.23 installed (check with `aws --version`).
```
AWS_DEFAULT_REGION=us-east-2; aws textract analyze-expense --document '{"S3Object": {"Bucket": "amazon-textract-public-content","Name": "blogs/textract-receipt-whole-foods-bryant-park.jpg"}}'
```

Python
The output is the Textract JSON response.
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/09/10/2-Textract-ReceiptAPI.jpg)
We also created a [Amazon Textract response parser library](https://github.com/aws-samples/amazon-textract-response-parser/blob/master/src-python/trp/trp2_expense.py) to parse the JSON returned by the AnalyzeExpense API. The library parses JSON and provides programming language-specific constructs to work with different parts of the document.
First install the dependencies.
```
> python3 -m pip install amazon-textract-response-parser boto3 amazon-textract-prettyprinter --upgrade
```

Python
This Python code takes the JSON response and prints out summary and line items in a table structure:
```
import os
import boto3
from textractprettyprinter.t_pretty_print_expense import get_string, Textract_Expense_Pretty_Print, Pretty_Print_Table_Format
textract = boto3.client(service_name='textract')
try:
  response = textract.analyze_expense(
    Document={
      'S3Object': {
        'Bucket': "amazon-textract-public-content",
        'Name': " blogs/textract-receipt-whole-foods-bryant-park.jpg "
      }
    })
  pretty_printed_string = get_string(textract_json=response, output_type=[Textract_Expense_Pretty_Print.SUMMARY, Textract_Expense_Pretty_Print.LINEITEMGROUPS], table_format=Pretty_Print_Table_Format.fancy_grid)
    
  
  print(pretty_printed_string)  
except Exception as e_raise:
  print(e_raise)
  raise e_raise

```

Python
#### Output from code
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/09/10/3-Textract-ReceiptAPI.jpg)
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/09/10/4-Textract-ReceiptAPI.jpg)
More details and examples to the [AnalyzeExpense](https://docs.aws.amazon.com/textract/latest/dg/API_AnalyzeExpense.html) action can be found in the blog post [Announcing specialized support for extracting data from invoices and receipts using Amazon Textract](https://aws.amazon.com/blogs/machine-learning/announcing-expanded-support-for-extracting-data-from-invoices-and-receipts-using-amazon-textract/).
### Extract information from identity documents
Analyze ID helps you automatically extract information from identification documents such as driver’s licenses and passports. Using the following sample image we can use the `amazon-textract-caller` and the `amazon-textract-response-parser` to quickly extract the information from the document.
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/12/02/Jorge-Souza-DL.jpg)
First install the dependencies.
```
> python3 -m pip install amazon-textract-response-parser boto3 tabulate --upgrade
```

Python
`tabulate` is only used for visualization purposes in this example and is not necessary for automation.
This script calls the Analyze ID API and prints out the values in a tabular format.
```
import boto3
import trp.trp2_analyzeid as t2id
# call_textract_analyzeid is a wrapper method to which the S3 locations, local file location or bytes of the id document can be passed
from textractcaller import call_textract_analyzeid
# the bucket hosting the sample image is stored in us-east-2, so our region should be us-east-2
textract_client = boto3.client('textract', region_name='us-east-2')
j = call_textract_analyzeid(
    document_pages=[
  "s3://amazon-textract-public-content/analyzeid/driverlicense.png"],
    boto3_textract_client=textract_client)
# deserializing to Python class
doc: t2id.TAnalyzeIdDocument = t2id.TAnalyzeIdDocumentSchema().load(j)
# get the list of values as array
result = doc.get_values_as_list()
from tabulate import tabulate
print(tabulate([x[1:3] for x in result]))

```

Python
The output in this case is just the key and value pairs. Analyze ID also returns the confidence score and normalized values when available.
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/12/02/Output-DL.jpg)
### Multi-column detection and reading order
Traditional OCR solutions read left to right and don’t detect multiple columns, so they may generate incorrect reading order for multi-column documents. In addition to detecting text, Amazon Textract provides additional geometry information that you can use to detect multiple columns and print the text in reading order.
The following image is a two-column document. Similar to the earlier example, the image isn’t good quality, but Amazon Textract still performs well.
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/06/09/9-4078-extract.jpg)
The following example code processes the document with Amazon Textract and takes advantage of geometry information to print the text in reading order:
```
import boto3
# Document
s3BucketName = "amazon-textract-public-content"
documentName = "blogs/two-column-image.jpg"
# Amazon Textract client
textract = boto3.client('textract')
# Call Amazon Textract
response = textract.detect_document_text(
  Document={
    'S3Object': {
      'Bucket': s3BucketName,
      'Name': documentName
    }
  })
#print(response)
# Detect columns and print lines
columns = []
lines = []
for item in response["Blocks"]:
   if item["BlockType"] == "LINE":
    column_found=False
    for index, column in enumerate(columns):
      bbox_left = item["Geometry"]["BoundingBox"]["Left"]
      bbox_right = item["Geometry"]["BoundingBox"]["Left"] + item["Geometry"]["BoundingBox"]["Width"]
      bbox_centre = item["Geometry"]["BoundingBox"]["Left"] + item["Geometry"]["BoundingBox"]["Width"]/2
      column_centre = column['left'] + column['right']/2
      if (bbox_centre > column['left'] and bbox_centre < column['right']) or (column_centre > bbox_left and column_centre < bbox_right):
        #Bbox appears inside the column
        lines.append([index, item["Text"]])
        column_found=True
        break
    if not column_found:
      columns.append({'left':item["Geometry"]["BoundingBox"]["Left"], 'right':item["Geometry"]["BoundingBox"]["Left"] + item["Geometry"]["BoundingBox"]["Width"]})
      lines.append([len(columns)-1, item["Text"]])
lines.sort(key=lambda x: x[0])
for line in lines:
  print (line[1])

```

Python
The following image shows the output of the detected text in the correct reading order.
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/06/09/10-4078-output.jpg)
## Natural language processing and document classification
Customer emails, support tickets, product reviews, social media, even advertising copy all represent insights into customer sentiment that can be put to work for your business. A lot of such content contains images or scanned versions of documents. After text is extracted from these documents, you can use [Amazon Comprehend](https://aws.amazon.com/comprehend/) to detect sentiment, entities, key phrases, syntax, and topics. You can also train Amazon Comprehend to detect custom entities based on your business domain. You can then use these insights to classify documents, automate business process workflows, and compliance.
The following example code processes the first image sample we used earlier with Amazon Textract to extract text and then uses Amazon Comprehend to detect sentiment and entities:
```
import boto3
# Document
s3BucketName = "amazon-textract-public-content"
documentName = "blogs/simple-document-image.jpg"
# Amazon Textract client
textract = boto3.client('textract')
# Call Amazon Textract
response = textract.detect_document_text(
  Document={
    'S3Object': {
      'Bucket': s3BucketName,
      'Name': documentName
    }
  })
#print(response)
# Print text
print("\nText\n========")
text = ""
for item in response["Blocks"]:
  if item["BlockType"] == "LINE":
    print ('\033[94m' + item["Text"] + '\033[0m')
    text = text + " " + item["Text"]
# Amazon Comprehend client
comprehend = boto3.client('comprehend')
# Detect sentiment
sentiment = comprehend.detect_sentiment(LanguageCode="en", Text=text)
print ("\nSentiment\n========\n{}".format(sentiment.get('Sentiment')))
# Detect entities
entities = comprehend.detect_entities(LanguageCode="en", Text=text)
print("\nEntities\n========")
for entity in entities["Entities"]:
  print ("{}\t=>\t{}".format(entity["Type"], entity["Text"]))

```

Python
The following image shows the output text along with the text analysis from Amazon Comprehend. It found the sentiment to be neutral and detected “Amazon” as an organization, “Seattle, WA” as a location, and “July 5th, 1994” as a date, along with other entities.
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/06/09/11-4078-text.jpg)
## Natural language processing for medical documents
An important way to improve patient care and accelerate clinical research is by understanding and analyzing the insights and relationships that are “trapped” in free-form medical text. These can include hospital admission notes and patient medical history.
In this example, we use the following document to extract text using Amazon Textract. You then use [Amazon Comprehend Medical](https://aws.amazon.com/comprehend/medical/) to extract medical entities, such as medical condition, medication, dosage, strength, and protected health information (PHI).
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/06/09/12-4078-patient-visit.jpg)
The following example code detects different medical entities:
```
import boto3
# Document
s3BucketName = "amazon-textract-public-content"
documentName = "blogs/medical-notes.png"
# Amazon Textract client
textract = boto3.client('textract')
# Call Amazon Textract
response = textract.detect_document_text(
  Document={
    'S3Object': {
      'Bucket': s3BucketName,
      'Name': documentName
    }
  })
#print(response)
# Print text
print("\nText\n========")
text = ""
for item in response["Blocks"]:
  if item["BlockType"] == "LINE":
    print ('\033[94m' + item["Text"] + '\033[0m')
    text = text + " " + item["Text"]
# Amazon Comprehend client
comprehend = boto3.client('comprehendmedical')
# Detect medical entities
entities = comprehend.detect_entities(Text=text)
print("\nMidical Entities\n========")
for entity in entities["Entities"]:
  print("- {}".format(entity["Text"]))
  print ("  Type: {}".format(entity["Type"]))
  print ("  Category: {}".format(entity["Category"]))
  if(entity["Traits"]):
    print("  Traits:")
    for trait in entity["Traits"]:
      print ("  - {}".format(trait["Name"]))
  print("\n")

```

Python
The following image and text block shows the output of the detected text with information categorized by type. It detected “40yo” as the age with category `Protected Health Information`. It also detected different medical conditions, including sleeping trouble, rash, inferior turbinates, and erythematous eruption. It recognized different medications and anatomy information.
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/06/09/13-4078-output.jpg)
```
Medical Entities
========
- 40yo
  Type: AGE
  Category: PROTECTED_HEALTH_INFORMATION
- Sleeping trouble
  Type: DX_NAME
  Category: MEDICAL_CONDITION
  Traits:
  - SYMPTOM
- Clonidine
  Type: GENERIC_NAME
  Category: MEDICATION
- Rash
  Type: DX_NAME
  Category: MEDICAL_CONDITION
  Traits:
  - SYMPTOM
- face
  Type: SYSTEM_ORGAN_SITE
  Category: ANATOMY
- leg
  Type: SYSTEM_ORGAN_SITE
  Category: ANATOMY
- Vyvanse
  Type: BRAND_NAME
  Category: MEDICATION
- Clonidine
  Type: GENERIC_NAME
  Category: MEDICATION
- HEENT
  Type: SYSTEM_ORGAN_SITE
  Category: ANATOMY
- Boggy inferior turbinates
  Type: DX_NAME
  Category: MEDICAL_CONDITION
  Traits:
  - SIGN
- inferior
  Type: DIRECTION
  Category: ANATOMY
- turbinates
  Type: SYSTEM_ORGAN_SITE
  Category: ANATOMY
- oropharyngeal lesion
  Type: DX_NAME
  Category: MEDICAL_CONDITION
  Traits:
  - SIGN
  - NEGATION
- Lungs
  Type: SYSTEM_ORGAN_SITE
  Category: ANATOMY
- clear Heart
  Type: DX_NAME
  Category: MEDICAL_CONDITION
  Traits:
  - SIGN
- Heart
  Type: SYSTEM_ORGAN_SITE
  Category: ANATOMY
- Regular rhythm
  Type: DX_NAME
  Category: MEDICAL_CONDITION
  Traits:
  - SIGN
- Skin
  Type: SYSTEM_ORGAN_SITE
  Category: ANATOMY
- erythematous eruption
  Type: DX_NAME
  Category: MEDICAL_CONDITION
  Traits:
  - SIGN
- hairline
  Type: SYSTEM_ORGAN_SITE
  Category: ANATOMY

```

Code
## Document translation
Many organizations localize content for international users, such as websites and applications. They must translate large volumes of documents efficiently. You can use Amazon Textract with [Amazon Translate](https://aws.amazon.com/translate/) to extract text and data and then translate them into other languages.
The following code example shows translating the text in the first image to German:
```
import boto3
# Document
s3BucketName = "amazon-textract-public-content"
documentName = "blogs/simple-document-image.jpg"
# Amazon Textract client
textract = boto3.client('textract')
# Call Amazon Textract
response = textract.detect_document_text(
  Document={
    'S3Object': {
      'Bucket': s3BucketName,
      'Name': documentName
    }
  })
#print(response)
# Amazon Translate client
translate = boto3.client('translate')
print ('')
for item in response["Blocks"]:
  if item["BlockType"] == "LINE":
    print ('\033[94m' + item["Text"] + '\033[0m')
    result = translate.translate_text(Text=item["Text"], SourceLanguageCode="en", TargetLanguageCode="de")
    print ('\033[92m' + result.get('TranslatedText') + '\033[0m')
  print ('')

```

Python
The following image shows the output of the detected text, translated to German line by line.
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/06/09/14-4078-output.jpg)
## Search and discovery
Extracting structured data from documents and creating a smart index using [Amazon OpenSearch Service](https://aws.amazon.com/elasticsearch-service/) allows you to search through millions of documents quickly. For example, a mortgage company could use Amazon Textract to process millions of scanned loan applications in a matter of hours and have the extracted data indexed in Amazon ES. This would allow them to create search experiences like searching for loan applications where the applicant’s name is John Doe, or searching for contracts where the interest rate is 2%.
The following code example extracts text from the first image, stores it in Amazon ES, and searches it using Kibana:
```
import boto3
from elasticsearch import Elasticsearch, RequestsHttpConnection
from requests_aws4auth import AWS4Auth
def indexDocument(bucketName, objectName, text):
  # Update host with endpoint of your Elasticsearch cluster
  #host = "search--xxxxxxxxxxxxxx.us-east-1.es.amazonaws.com
  host = "searchxxxxxxxxxxxxxxxx.us-east-1.es.amazonaws.com"
  region = 'us-east-1'
  if(text):
    service = 'es'
    ss = boto3.Session()
    credentials = ss.get_credentials()
    region = ss.region_name
    awsauth = AWS4Auth(credentials.access_key, credentials.secret_key, region, service, session_token=credentials.token)
    es = Elasticsearch(
      hosts = [{'host': host, 'port': 443}],
      http_auth = awsauth,
      use_ssl = True,
      verify_certs = True,
      connection_class = RequestsHttpConnection
    )
    document = {
      "name": "{}".format(objectName),
      "bucket" : "{}".format(bucketName),
      "content" : text
    }
    es.index(index="textract", doc_type="document", id=objectName, body=document)
    print("Indexed document: {}".format(objectName))
# Document
s3BucketName = "amazon-textract-public-content"
documentName = "blogs/simple-document-image.jpg"
# Amazon Textract client
textract = boto3.client('textract')
# Call Amazon Textract
response = textract.detect_document_text(
  Document={
    'S3Object': {
      'Bucket': s3BucketName,
      'Name': documentName
    }
  })
#print(response)
# Print detected text
text = ""
for item in response["Blocks"]:
  if item["BlockType"] == "LINE":
    print ('\033[94m' + item["Text"] + '\033[0m')
    text += item["Text"]
indexDocument(s3BucketName, documentName, text)
# You can view index documents in Kibana Dashboard

```

Python
The following image shows the output of extracted text in Kibana search results.
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/06/09/15-4078-kibana.jpg)
You can also build a custom UI experience by taking advantage of the Amazon ES APIs. Later in the post, you learn how to extract forms and tables and then index that structured data similarly to enable smart search.
## Compliance control with document redaction
Because Amazon Textract identifies data types and form labels automatically, AWS helps secure infrastructure so that you can maintain compliance with information controls. For example, an insurer could use Amazon Textract to feed a workflow that automatically redacts personally identifiable information (PII) for review before archiving claim forms. Amazon Textract recognizes the important fields that require protection.
The following code example form fields in the employment application used earlier and address fields:
```
import boto3
from trp import Document
from PIL import Image, ImageDraw
# Document
s3BucketName = "amazon-textract-public-content"
documentName = "blogs/employeeapp20210510.png"
# Amazon Textract client
textract = boto3.client('textract')
# Call Amazon Textract
response = textract.analyze_document(
  Document={
    'S3Object': {
      'Bucket': s3BucketName,
      'Name': documentName
    }
  },
  FeatureTypes=["FORMS"])
doc = Document(response)
# Redact document
img = Image.open(documentName)
width, height = img.size
if(doc.pages):
  page = doc.pages[0]
  for field in page.form.fields:
    if(field.key and field.value and "address" in field.key.text.lower()):
    #if(field.key and field.value):
      print("Redacting => Key: {}, Value: {}".format(field.key.text, field.value.text))
      
      x1 = field.value.geometry.boundingBox.left*width
      y1 = field.value.geometry.boundingBox.top*height-2
      x2 = x1 + (field.value.geometry.boundingBox.width*width)+5
      y2 = y1 + (field.value.geometry.boundingBox.height*height)+2
      draw = ImageDraw.Draw(img)
      draw.rectangle([x1, y1, x2, y2], fill="Black")
img.save("redacted-{}".format(documentName))  

```

Python
The following output is the redacted version of the employment application.
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/06/09/16-4078-application.jpg)
## PDF or multi-page TIFF document processing (asynchronous API operations)
For the earlier examples, you used images with the synchronous API operations. Now, see how we process PDF files using the asynchronous API operations. Single page or multi-page TIFF documents are also supported in the asynchronous API operations.
With the `amazon-textract` command line tool, you can pass in a PDF (the location for the PDF has to be on Amazon S3) and the underlying implementation calls the asynchronous API for [StartDocumentTextDetection](https://docs.aws.amazon.com/textract/latest/dg/API_StartDocumentTextDetection.html) or [StartDocumentAnalysis](https://docs.aws.amazon.com/textract/latest/dg/API_StartDocumentAnalysis.html) to start an Amazon Textract job:
```
amazon-textract --input-document "s3://amazon-textract-public-content/blogs/Amazon-Textract-Pdf.pdf" --pretty-print LINES
```

Code
The following screenshot shows our output.
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/06/09/17-4078-screenshot.jpg)
When you use the asynchronous API from a Python program or the Python Interpreter, it looks like the following code:
```
from textractcaller.t_call import call_textract
from textractprettyprinter.t_pretty_print import get_lines_string
response = call_textract(input_document="s3://amazon-textract-public-content/blogs/Amazon-Textract-Pdf.pdf")
print(get_lines_string(response))

```

Python
We get the following output.
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/06/09/18-4078-screenshot.jpg)
First, [StartDocumentTextDetection](https://docs.aws.amazon.com/textract/latest/dg/API_StartDocumentTextDetection.html) or [StartDocumentAnalysis](https://docs.aws.amazon.com/textract/latest/dg/API_StartDocumentAnalysis.html) is called to start an Amazon Textract job. Amazon Textract publishes the results of the Amazon Textract request, including completion status, to [Amazon Simple Notification Service](https://aws.amazon.com/sns) (Amazon SNS). You can then use [GetDocumentTextDetection](https://docs.aws.amazon.com/textract/latest/dg/API_GetDocumentTextDetection.html) or [GetDocumentAnalysis](https://docs.aws.amazon.com/textract/latest/dg/API_GetDocumentAnalysis.html) to get the results from Amazon Textract.
## Conclusion
In this post, we showed you how to use [Amazon Textract](https://aws.amazon.com/textract) to automatically extract text and data from scanned documents without ML experience. We covered use cases in fields such as finance, healthcare, and HR, but there are many other opportunities in which the ability to unlock text and data from unstructured documents could be useful.
You can start using Amazon Textract in the Regions US East (Ohio), US East (Northern Virginia), US West (N. California), US West (Oregon), Asia Pacific (Mumbai), Asia Pacific (Seoul), Asia Pacific (Singapore), Asia Pacific (Sydney), Canada (Central), EU (Frankfurt), EU (Ireland), EU (London), EU (Paris), AWS GovCloud (US-East), AWS GovCloud (US-West) and Europe (Spain).
To learn more about Amazon Textract, read about processing [single-page](https://docs.aws.amazon.com/textract/latest/dg/sync.html) and [multipage](https://docs.aws.amazon.com/textract/latest/dg/async.html) documents, [working with block objects](https://docs.aws.amazon.com/textract/latest/dg/how-it-works-document-layout.html), and [code samples](https://github.com/aws-samples/amazon-textract-code-samples).
### About the Authors
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2018/04/03/kashif-imran-100.jpg)**Kashif Imran** is a Solutions Architect at Amazon Web Services. He works with some of the largest strategic AWS customers to provide technical guidance and design advice. His expertise spans application architecture, serverless, containers, NoSQL and machine learning.
**![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2021/06/10/Martin-Schade.jpg)Martin Schade** is a Senior ML Product SA with the Amazon Textract team. He has 20+ years of experience with internet-related technologies, engineering and architecting solutions and joined AWS in 2014, first guiding some of the largest AWS customers on most efficient and scalable use of AWS services and later focused on AI/ML with a focus on computer vision and at the moment is obsessed with extracting information from documents.
![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2025/06/02/urvi.jpg)**Urvi Sharma** is a Solutions Architect at AWS who is passionate about working on edge services. She works with customers in the early stages of cloud adoption to help them migrate and modernize, and build resilient and secure architectures and incorporate AI/ML services with modern technologies like generative AI.
**![](https://d2908q01vomqb2.cloudfront.net/f1f836cb4ea6efb2a0b1b99f41ad8b103eff4b59/2025/06/02/sameer.jpg)Sameer Shrivastava** is a Solutions Architect at AWS, where he specializes in serverless technologies. He is enthusiastic about guiding organizations in their cloud transformation journeys to drive efficiency and scalability.
[ ![](https://d1.awsstatic.com/Digital%20Marketing/House/Editorial/other/SiteMerch-3066-Podcast_Editorial.65839609a8dda387937ed07dc8dc4f3c3b870546.png) AWS Podcast  Subscribe for weekly AWS news and interviews  Learn more  ](https://aws.amazon.com/podcasts/aws-podcast/?sc_icampaign=aware_aws-podcast&sc_ichannel=ha&sc_icontent=awssm-2021&sc_iplace=blog_tile&trk=ha_awssm-2021)
[ ![](https://d1.awsstatic.com/webteam/homepage/editorials/Site-Merch_APN_Editorial.12df33fb7e0299389b086fb48dba7b9deeef07df.png) AWS Partner Network  Find an APN member to support your cloud business needs  Learn more  ](https://aws.amazon.com/partners/find/?sc_icampaign=aware_apn_recruit&sc_ichannel=ha&sc_icontent=awssm-2021&sc_iplace=blog_tile&trk=ha_awssm-2021)
[ ![](https://d1.awsstatic.com/webteam/homepage/editorials/Site-Merch_Training_Editorial.5cc72ab0552ba66ef4e36a1a60ee742bc31113c7.png) AWS Training & Certifications  Free digital courses to help you develop your skills  Learn more  ](https://aws.amazon.com/training/?sc_icampaign=aware_aws-training_blog&sc_ichannel=ha&sc_icontent=awssm-2021&sc_iplace=blog_tile&trk=ha_awssm-2021)
###  Resources
  * [Getting Started](https://aws.amazon.com/getting-started?sc_ichannel=ha&sc_icampaign=acq_awsblogsb&sc_icontent=machine-learning-resources)
  * [What's New](https://aws.amazon.com/new?sc_ichannel=ha&sc_icampaign=acq_awsblogsb&sc_icontent=machine-learning-resources)


###  Blog Topics
  * [Amazon Bedrock](https://aws.amazon.com/blogs/machine-learning/category/artificial-intelligence/amazon-machine-learning/amazon-bedrock/)
  * [Amazon Comprehend](https://aws.amazon.com/blogs/machine-learning/category/artificial-intelligence/amazon-comprehend/)
  * [Amazon Kendra](https://aws.amazon.com/blogs/machine-learning/category/artificial-intelligence/amazon-kendra/)
  * [Amazon Lex](https://aws.amazon.com/blogs/machine-learning/category/artificial-intelligence/amazon-lex/)
  * [Amazon Polly](https://aws.amazon.com/blogs/machine-learning/category/artificial-intelligence/amazon-polly/)
  * [Amazon Q](https://aws.amazon.com/blogs/machine-learning/category/amazon-q/)
  * [Amazon Rekognition](https://aws.amazon.com/blogs/machine-learning/category/artificial-intelligence/amazon-rekognition/)
  * [Amazon SageMaker](https://aws.amazon.com/blogs/machine-learning/category/artificial-intelligence/sagemaker/)
  * [Amazon Textract](https://aws.amazon.com/blogs/machine-learning/category/artificial-intelligence/amazon-textract/)


###  Follow
  * [ Twitter](https://twitter.com/awscloud)
  * [ Facebook](https://www.facebook.com/amazonwebservices)
  * [ LinkedIn](https://www.linkedin.com/company/amazon-web-services/)
  * [ Twitch](https://www.twitch.tv/aws)
  * [ Email Updates](https://pages.awscloud.com/communication-preferences?sc_ichannel=ha&sc_icampaign=acq_awsblogsb&sc_icontent=maching-learning-social)


[Create an AWS account](https://portal.aws.amazon.com/gp/aws/developer/registration/index.html?nc1=f_ct&src=footer_signup)
## Learn
  * [What Is AWS?](https://aws.amazon.com/what-is-aws/?nc1=f_cc)
  * [What Is Cloud Computing?](https://aws.amazon.com/what-is-cloud-computing/?nc1=f_cc)
  * [What Is Generative AI?](https://aws.amazon.com/what-is/generative-ai/?nc1=f_cc)
  * [Cloud Computing Concepts Hub](https://aws.amazon.com/what-is/?nc1=f_cc)
  * [AWS Cloud Security](https://aws.amazon.com/security/?nc1=f_cc)
  * [What's New](https://aws.amazon.com/new/?nc1=f_cc)
  * [Blogs](https://aws.amazon.com/blogs/?nc1=f_cc)
  * [Press Releases](https://press.aboutamazon.com/press-releases/aws)


## Resources
  * [Getting Started](https://aws.amazon.com/getting-started/?nc1=f_cc)
  * [Training](https://aws.amazon.com/training/?nc1=f_cc)
  * [AWS Solutions Library](https://aws.amazon.com/solutions/?nc1=f_cc)
  * [Architecture Center](https://aws.amazon.com/architecture/?nc1=f_cc)
  * [Product and Technical FAQs](https://aws.amazon.com/faqs/?nc1=f_dr)
  * [Analyst Reports](https://aws.amazon.com/resources/analyst-reports/?nc1=f_cc)
  * [AWS Partners](https://aws.amazon.com/partners/work-with-partners/?nc1=f_dr)
  * [AWS Inclusion, Diversity & Equity](https://aws.amazon.com/diversity-inclusion/?nc1=f_cc)


## Developers
  * [Developer Center](https://aws.amazon.com/developer/?nc1=f_dr)
  * [SDKs & Tools](https://aws.amazon.com/developer/tools/?nc1=f_dr)
  * [.NET on AWS](https://aws.amazon.com/developer/language/net/?nc1=f_dr)
  * [Python on AWS](https://aws.amazon.com/developer/language/python/?nc1=f_dr)
  * [Java on AWS](https://aws.amazon.com/developer/language/java/?nc1=f_dr)
  * [PHP on AWS](https://aws.amazon.com/developer/language/php/?nc1=f_cc)
  * [JavaScript on AWS](https://aws.amazon.com/developer/language/javascript/?nc1=f_dr)


## Help
  * [Contact Us](https://aws.amazon.com/contact-us/?nc1=f_m)
  * [File a Support Ticket](https://console.aws.amazon.com/support/home/?nc1=f_dr)
  * [AWS re:Post](https://repost.aws/?nc1=f_dr)
  * [Knowledge Center](https://repost.aws/knowledge-center/?nc1=f_dr)
  * [AWS Support Overview](https://aws.amazon.com/premiumsupport/?nc1=f_dr)
  * [Get Expert Help](https://iq.aws.amazon.com/?utm=mkt.foot/?nc1=f_m)
  * [AWS Accessibility](https://aws.amazon.com/accessibility/?nc1=f_cc)
  * [Legal](https://aws.amazon.com/legal/?nc1=f_cc)


English
Back to top
Amazon is an Equal Opportunity Employer: Minority / Women / Disability / Veteran / Gender Identity / Sexual Orientation / Age.
[](https://twitter.com/awscloud)[facebook](https://www.facebook.com/amazonwebservices)[linkedin](https://www.linkedin.com/company/amazon-web-services/)[instagram](https://www.instagram.com/amazonwebservices/)[twitch](https://www.twitch.tv/aws)[youtube](https://www.youtube.com/user/AmazonWebServices/Cloud/)[podcasts](https://aws.amazon.com/podcasts/?nc1=f_cc)[email](https://pages.awscloud.com/communication-preferences?trk=homepage)
  * [Privacy](https://aws.amazon.com/privacy/?nc1=f_pr)
  * [Site terms](https://aws.amazon.com/terms/?nc1=f_pr)
  * [Cookie Preferences](https://aws.amazon.com/blogs/machine-learning/automatically-extract-text-and-structured-data-from-documents-with-amazon-textract/)


© 2025, Amazon Web Services, Inc. or its affiliates. All rights reserved.


## Source Information
- URL: https://aws.amazon.com/blogs/machine-learning/automatically-extract-text-and-structured-data-from-documents-with-amazon-textract/
- Harvested: 2025-08-19T20:31:12.908961
- Profile: code_examples
- Agent: architect
