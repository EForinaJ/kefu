<template>
    <div class="art-card p-4 box-border mb-5 max-sm:mb-4">
        <div class="art-card-header mb-4">
            <div class="title">
                <h4>常用功能</h4>
            </div>
        </div>
        <ElRow :gutter="20">
            <ElCol :sm="24" :md="12" :lg="6">
                <div class="flex flex-col items-center justify-center cursor-pointer">
                    <div  class="flex-cc mb-3 size-10 rounded-lg bg-theme/12 text-theme">
                        <ArtSvgIcon icon="ri:shopping-bag-4-line" class="text-xl" />
                    </div>
                    <div class="text-sm">快速下单</div>
                </div>
            </ElCol>
            <ElCol :sm="24" :md="12" :lg="6">
                <div class="flex flex-col items-center justify-center cursor-pointer">
                    <div  class="flex-cc mb-3 size-10 rounded-lg bg-theme/12 text-theme">
                        <ArtSvgIcon icon="ri:shopping-bag-4-line" class="text-xl" />
                    </div>
                    <div class="text-sm">售后处理</div>
                </div>
            </ElCol>
            <ElCol :sm="24" :md="12" :lg="6">
                <div class="flex flex-col items-center justify-center cursor-pointer">
                    <div  class="flex-cc mb-3 size-10 rounded-lg bg-theme/12 text-theme">
                        <ArtSvgIcon icon="ri:shopping-bag-4-line" class="text-xl" />
                    </div>
                    <div class="text-sm">添加威客</div>
                </div>
            </ElCol>
            <!-- <ElCol :sm="24" :md="12" :lg="6">
                <div class="flex flex-col items-center justify-center">
                    <div  class="flex-cc mb-3 size-10 rounded-lg bg-theme/12 text-theme">
                        <ArtSvgIcon icon="ri:shopping-bag-4-line" class="text-xl" />
                    </div>
                    <div class="text-sm">添加工单</div>
                </div>
            </ElCol> -->
        </ElRow>
    </div>
</template>