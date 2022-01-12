package main

var (
	label_enter_project                   = "✏️ ثبت آگهی پروژه"
	label_home                            = "صفحه اصلی 🏠"
	label_cancel_entring_project_proccess = "❌ از ثبت آگهی منصرف شدم"
	label_accept                          = "تایید"
	label_reject                          = "رد"
	label_project_accepted                = "پروژه تایید شد✅"
	label_project_entered                 = "✅ پروژه ثبت گردید. پس از تایید انتشار میابد."
	label_project_rejected                = "پروژه رد شد🚫"
	label_message_to_owner                = "پیام به کارفرما"
	label_project_status_connot_edit      = "وضعیت پروژه %v قبلا مشخص شده و قابل تغییر نیست."
	label_message_enter_same_project      = "ثبت پروژه مشابه"
	label_please_enter_poster_title       = `
	👇 موضوع آگهی خود را ارسال کنید
برای مثال: طراحی وب سایت

توضیحات آگهی شما در بخش بعدی گرفته میشود، لطفا از درج توضیحات در این بخش خودداری فرمایید.
`
	label_please_enter_poster_description = `✍️  لطفا متن آگهی خود را وارد کنید.
	برای مثال: به یک نفر جهت طراحی سایت فروشگاهی نیازمندیم.

	نکات مهم:
	❗️ اگر متن اگهی قوانین کانال را نقض کند آگهی رد میشود
	❗️ اگر متن آگهی قوانین جمهوری اسلامی را نقض کند ، آگهی رد میشود.
	❗️ نوشتن شماره تلفن ، آی دی و... در متن آگهی ممنوع است.
	`
	label_help                   = "💡 راهنما "
	label_support                = "☎️ پشتیبانی"
	label_bot_designer           = "💻 طراح ربات"
	label_users_list             = "لیست کاربران"
	label_accept_bot_usage_roles = "✅ قوانین را میپذیرم"
)
var (
	description_must_have_id = `
	⚠️ برای ثبت آگهی ابتدا باید برای خودت آیدی بسازی ⚠️
	از طریق تنظیمات تلگرام خود میتوانید اینکار را انجام دهید
	
	بعد از ساختن آیدی به ربات برگردید و آگهی خود را ثبت کنید`
	description_should_have_phone_number = `
	⚠️ برای ثبت آگهی در ربات، ابتدا میبایست شماره موبایل خود را برای ربات ارسال کنید؛ با استفاده از دکمه زیر میتوانید شماره خود را ارسال کنید`
	description_command_not_found = "دستور پیدا نشد! \n لطفا یکی از گزینه های زیر را انتخاب کنید."
	description_project_poster    = `
	• %v
	%v
	@getprojectsofficial
	➖➖➖➖➖➖➖
	برای ثبت آگهی به ربات @getprojectsbot مراجعه کنید.
`
	description_project_accepted      = "✅ پروژه %v تایید شد و در کانال قرار گرفت."
	description_project_rejected      = "🚫 پروژه  %v رد شد."
	description_new_project_from_user = `🚧 کاربر %v پروژه ای با عنوان « %v »
	و متن:%v
	 ثبت کرده است.
	`
	description_project_entering_canceled = `
	ثبت پروژه کنسل شد.
	برای ادامه از دکمه های زیر استفاده کنید:`
	description_help = `
	سوالات متداول و راهنمای کانال پراجکت 💡

⁉️ چطوری آگهی ثبت کنم؟!
https://t.me/prajects_help/9
`
	description_support = `
برای ثبت پروژه به کمک نیاز دارید؟
با ادمین کانال در ارتباط باشید:
@GetProjectsAdmin
➖➖➖➖➖➖➖
@getprojectsofficial
`
	description_bot_designer = `ارائه انواع خدمات بصورت حرفه ای:

🔸طراحی قالب و وب سایت
🔹طراحی ربات تلگرام
🔸انجام پروژه های دانشجویی با زبان: 
HTML, CSS, JS, C# , Go , C++

ارتباط با ما:
@Moh3n_dev`
	description_bot_usage_roles = `
❗ کانال در قبال اگهی دهنده و محتوای مورد نظر شخص اگهی دهنده هیچ مسئولیتی ندارد.
و فقط هدف این کانال درج اگهی و تبلیغات هست

❗ درصورت گزارش تخلف، اگهی و فرد متخلف فورا از کانال حذف میشوند`
	log_text            = "username : %v user-id : %v text : %v first-name : %v last-name : %v bio : %v cache : %v"
	description_welcome = `سلام %v 🖐
	به  کانال Get Projects خوش اومدی!
	🔹 اينجا تو ميتونی به راحتي آگهیت رو درج كنی، فقط کافیه دکمه ثبت آگهی رو انتخاب بکنی
	
	🔸 اگر انتقاد، نظر و یا پیشنهادی داشتید میتونید با آیدی زیر در میون بزارید
	@GetProjectsAdmin`
)
